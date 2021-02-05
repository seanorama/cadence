// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history

import (
	"sync"
	"time"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/backoff"
	"github.com/uber/cadence/common/cache"
	"github.com/uber/cadence/common/clock"
	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/log/tag"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/persistence"
	"github.com/uber/cadence/common/types"
	"github.com/uber/cadence/service/history/config"
	"github.com/uber/cadence/service/history/execution"
	"github.com/uber/cadence/service/history/shard"
	"github.com/uber/cadence/service/history/task"
)

const (
	firstPhaseRetryInitialDelay  = 50 * time.Millisecond
	firstPhaseRetryCount         = 3
	secondPhaseRetryInitialDelay = 2 * time.Second
	secondPhaseRetryMaxDelay     = 128 * time.Second
	secondPhaseRetryExpiration   = 5 * time.Minute
)

type (
	taskProcessorOptions struct {
		queueSize   int
		workerCount int
	}

	taskInfo struct {
		// TODO: change to queueTaskExecutor
		processor taskExecutor
		task      task.Info

		attempt   int
		startTime time.Time
		logger    log.Logger

		// used by 2DC task life cycle
		// TODO remove when NDC task life cycle is implemented
		shouldProcessTask bool
	}

	taskProcessor struct {
		shard                   shard.Context
		shutdownCh              chan struct{}
		tasksCh                 chan *taskInfo
		config                  *config.Config
		logger                  log.Logger
		metricsClient           metrics.Client
		timeSource              clock.TimeSource
		retryPolicy             backoff.RetryPolicy
		workerWG                sync.WaitGroup
		domainMetricsScopeCache cache.DomainMetricsScopeCache

		// worker coroutines notification
		workerNotificationChans []chan struct{}
		// duplicate numOfWorker from config.TimerTaskWorkerCount for dynamic config works correctly
		numOfWorker int
	}
)

func newTaskInfo(
	processor taskExecutor,
	task task.Info,
	logger log.Logger,
	startTime time.Time,
) *taskInfo {
	return &taskInfo{
		processor:         processor,
		task:              task,
		attempt:           0,
		startTime:         startTime, // used for metrics
		logger:            logger,
		shouldProcessTask: true,
	}
}

func newTaskProcessor(
	options taskProcessorOptions,
	shard shard.Context,
	executionCache *execution.Cache,
	logger log.Logger,
) *taskProcessor {

	workerNotificationChans := []chan struct{}{}
	for index := 0; index < options.workerCount; index++ {
		workerNotificationChans = append(workerNotificationChans, make(chan struct{}, 1))
	}
	firstPolicy := backoff.NewExponentialRetryPolicy(firstPhaseRetryInitialDelay)
	firstPolicy.SetMaximumAttempts(firstPhaseRetryCount)
	secondPolicy := backoff.NewExponentialRetryPolicy(secondPhaseRetryInitialDelay)
	secondPolicy.SetMaximumInterval(secondPhaseRetryMaxDelay)
	secondPolicy.SetExpirationInterval(secondPhaseRetryExpiration)
	base := &taskProcessor{
		shard:                   shard,
		shutdownCh:              make(chan struct{}),
		tasksCh:                 make(chan *taskInfo, options.queueSize),
		config:                  shard.GetConfig(),
		logger:                  logger,
		metricsClient:           shard.GetMetricsClient(),
		domainMetricsScopeCache: shard.GetService().GetDomainMetricsScopeCache(),
		timeSource:              shard.GetTimeSource(),
		workerNotificationChans: workerNotificationChans,
		retryPolicy:             backoff.NewMultiPhasesRetryPolicy(firstPolicy, secondPolicy),
		numOfWorker:             options.workerCount,
	}

	return base
}

func (t *taskProcessor) start() {
	for i := 0; i < t.numOfWorker; i++ {
		t.workerWG.Add(1)
		notificationChan := t.workerNotificationChans[i]
		go t.taskWorker(notificationChan)
	}
	t.logger.Info("Task processor started.")
}

func (t *taskProcessor) stop() {
	close(t.shutdownCh)
	if success := common.AwaitWaitGroup(&t.workerWG, time.Minute); !success {
		t.logger.Warn("Task processor timed out on shutdown.")
	}
	t.logger.Debug("Task processor shutdown.")
}

func (t *taskProcessor) taskWorker(
	notificationChan chan struct{},
) {
	defer t.workerWG.Done()

	for {
		select {
		case <-t.shutdownCh:
			return
		case task, ok := <-t.tasksCh:
			if !ok {
				return
			}
			t.processTaskAndAck(notificationChan, task)
		}
	}
}

func (t *taskProcessor) retryTasks() {
	for _, workerNotificationChan := range t.workerNotificationChans {
		select {
		case workerNotificationChan <- struct{}{}:
		default:
		}
	}
}

func (t *taskProcessor) addTask(
	task *taskInfo,
) bool {
	// We have a timer to fire.
	select {
	case t.tasksCh <- task:
	case <-t.shutdownCh:
		return false
	}
	return true
}

func (t *taskProcessor) processTaskAndAck(
	notificationChan <-chan struct{},
	task *taskInfo,
) {

	var scope metrics.Scope
	var err error

FilterLoop:
	for {
		select {
		case <-t.shutdownCh:
			// this must return without ack
			return
		default:
			task.shouldProcessTask, err = task.processor.getTaskFilter()(task.task)
			if err == nil {
				break FilterLoop
			}
			time.Sleep(loadDomainEntryForTimerTaskRetryDelay)
		}
	}

	op := func() error {
		scope, err = t.processTaskOnce(notificationChan, task)
		err := t.handleTaskError(scope, task, notificationChan, err)
		if err != nil {
			task.attempt++
			if task.attempt >= t.config.TimerTaskMaxRetryCount() {
				scope.RecordTimer(metrics.TaskAttemptTimerPerDomain, time.Duration(task.attempt))
				task.logger.Error("Critical error processing task, retrying.",
					tag.Error(err), tag.OperationCritical, tag.TaskType(task.task.GetTaskType()))
			}
		}
		return err
	}
	retryCondition := func(err error) bool {
		select {
		case <-t.shutdownCh:
			return false
		default:
			return true
		}
	}

	for {
		select {
		case <-t.shutdownCh:
			// this must return without ack
			return
		default:
			err = backoff.Retry(op, t.retryPolicy, retryCondition)
			if err == nil {
				t.ackTaskOnce(scope, task)
				return
			}
		}
	}
}

func (t *taskProcessor) processTaskOnce(
	notificationChan <-chan struct{},
	taskInfo *taskInfo,
) (metrics.Scope, error) {

	select {
	case <-notificationChan:
	default:
	}

	var scopeIdx int
	var err error

	startTime := t.timeSource.Now()
	scopeIdx, err = taskInfo.processor.process(taskInfo)

	scope := task.GetOrCreateDomainTaggedScope(t.shard, scopeIdx, taskInfo.task.GetDomainID(), t.logger)
	if taskInfo.shouldProcessTask {
		scope.IncCounter(metrics.TaskRequestsPerDomain)
		scope.RecordTimer(metrics.TaskProcessingLatencyPerDomain, time.Since(startTime))
	}

	return scope, err
}

func (t *taskProcessor) handleTaskError(
	scope metrics.Scope,
	taskInfo *taskInfo,
	notificationChan <-chan struct{},
	err error,
) error {

	if err == nil {
		return nil
	}

	if _, ok := err.(*types.EntityNotExistsError); ok {
		return nil
	}

	if transferTask, ok := taskInfo.task.(*persistence.TransferTaskInfo); ok &&
		transferTask.TaskType == persistence.TransferTaskTypeCloseExecution &&
		err == execution.ErrMissingWorkflowStartEvent &&
		t.config.EnableDropStuckTaskByDomainID(taskInfo.task.GetDomainID()) { // use domainID here to avoid accessing domainCache
		scope.IncCounter(metrics.TransferTaskMissingEventCounterPerDomain)
		taskInfo.logger.Error("Drop close execution transfer task due to corrupted workflow history", tag.Error(err), tag.LifeCycleProcessingFailed)
		return nil
	}

	// this is a transient error
	if err == task.ErrTaskRedispatch {
		scope.IncCounter(metrics.TaskStandbyRetryCounterPerDomain)
		select {
		case <-notificationChan:
		case <-t.shutdownCh:
		}
		return err
	}

	// this is a transient error during graceful failover
	if err == task.ErrTaskPendingActive {
		scope.IncCounter(metrics.TaskPendingActiveCounterPerDomain)
		select {
		case <-notificationChan:
		case <-t.shutdownCh:
		}
		return err
	}

	if err == task.ErrTaskDiscarded {
		scope.IncCounter(metrics.TaskDiscardedPerDomain)
		err = nil
	}

	// this is a transient error
	// TODO remove this error check special case
	//  since the new task life cycle will not give up until task processed / verified
	if _, ok := err.(*types.DomainNotActiveError); ok {
		if t.timeSource.Now().Sub(taskInfo.startTime) > 2*cache.DomainCacheRefreshInterval {
			scope.IncCounter(metrics.TaskNotActiveCounterPerDomain)
			return nil
		}

		return err
	}

	scope.IncCounter(metrics.TaskFailuresPerDomain)

	if _, ok := err.(*persistence.CurrentWorkflowConditionFailedError); ok {
		taskInfo.logger.Error("More than 2 workflow are running.", tag.Error(err), tag.LifeCycleProcessingFailed)
		return nil
	}

	if taskInfo.attempt > t.config.TimerTaskMaxRetryCount() && common.IsStickyTaskConditionError(err) {
		// sticky task could end up into endless loop in rare cases and
		// cause worker to keep getting decision timeout unless restart.
		// return nil here to break the endless loop
		return nil
	}

	taskInfo.logger.Error("Fail to process task", tag.Error(err), tag.LifeCycleProcessingFailed)
	return err
}

func (t *taskProcessor) ackTaskOnce(
	scope metrics.Scope,
	task *taskInfo,
) {

	task.processor.complete(task)
	if task.shouldProcessTask {
		scope.RecordTimer(metrics.TaskAttemptTimerPerDomain, time.Duration(task.attempt))
		scope.RecordTimer(metrics.TaskLatencyPerDomain, time.Since(task.startTime))
		scope.RecordTimer(metrics.TaskQueueLatencyPerDomain, time.Since(task.task.GetVisibilityTimestamp()))
	}
}