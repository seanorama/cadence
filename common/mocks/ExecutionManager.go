// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	persistence "github.com/uber/cadence/common/persistence"
)

// ExecutionManager is an autogenerated mock type for the ExecutionManager type
type ExecutionManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ExecutionManager) Close() {
	_m.Called()
}

// CompleteCrossClusterTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CompleteCrossClusterTask(ctx context.Context, request *persistence.CompleteCrossClusterTaskRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CompleteCrossClusterTaskRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteReplicationTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CompleteReplicationTask(ctx context.Context, request *persistence.CompleteReplicationTaskRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CompleteReplicationTaskRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteTimerTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CompleteTimerTask(ctx context.Context, request *persistence.CompleteTimerTaskRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CompleteTimerTaskRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteTransferTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CompleteTransferTask(ctx context.Context, request *persistence.CompleteTransferTaskRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CompleteTransferTaskRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConflictResolveWorkflowExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) ConflictResolveWorkflowExecution(ctx context.Context, request *persistence.ConflictResolveWorkflowExecutionRequest) (*persistence.ConflictResolveWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.ConflictResolveWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.ConflictResolveWorkflowExecutionRequest) *persistence.ConflictResolveWorkflowExecutionResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ConflictResolveWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.ConflictResolveWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFailoverMarkerTasks provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CreateFailoverMarkerTasks(ctx context.Context, request *persistence.CreateFailoverMarkersRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CreateFailoverMarkersRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWorkflowExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CreateWorkflowExecution(ctx context.Context, request *persistence.CreateWorkflowExecutionRequest) (*persistence.CreateWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.CreateWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CreateWorkflowExecutionRequest) *persistence.CreateWorkflowExecutionResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CreateWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.CreateWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCurrentWorkflowExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) DeleteCurrentWorkflowExecution(ctx context.Context, request *persistence.DeleteCurrentWorkflowExecutionRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.DeleteCurrentWorkflowExecutionRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReplicationTaskFromDLQ provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) DeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.DeleteReplicationTaskFromDLQRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.DeleteReplicationTaskFromDLQRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWorkflowExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) DeleteWorkflowExecution(ctx context.Context, request *persistence.DeleteWorkflowExecutionRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.DeleteWorkflowExecutionRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCrossClusterTasks provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetCrossClusterTasks(ctx context.Context, request *persistence.GetCrossClusterTasksRequest) (*persistence.GetCrossClusterTasksResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetCrossClusterTasksResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetCrossClusterTasksRequest) *persistence.GetCrossClusterTasksResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetCrossClusterTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetCrossClusterTasksRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetCurrentExecution(ctx context.Context, request *persistence.GetCurrentExecutionRequest) (*persistence.GetCurrentExecutionResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetCurrentExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetCurrentExecutionRequest) *persistence.GetCurrentExecutionResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetCurrentExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetCurrentExecutionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetName provides a mock function with given fields:
func (_m *ExecutionManager) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetReplicationDLQSize provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetReplicationDLQSize(ctx context.Context, request *persistence.GetReplicationDLQSizeRequest) (*persistence.GetReplicationDLQSizeResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetReplicationDLQSizeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetReplicationDLQSizeRequest) *persistence.GetReplicationDLQSizeResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetReplicationDLQSizeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetReplicationDLQSizeRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationTasks provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetReplicationTasks(ctx context.Context, request *persistence.GetReplicationTasksRequest) (*persistence.GetReplicationTasksResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetReplicationTasksResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetReplicationTasksRequest) *persistence.GetReplicationTasksResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetReplicationTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetReplicationTasksRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountReplicationTasks provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) CountReplicationTasks(ctx context.Context, request *persistence.CountReplicationTasksRequest) (*persistence.CountReplicationTasksResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.CountReplicationTasksResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.CountReplicationTasksRequest) *persistence.CountReplicationTasksResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CountReplicationTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.CountReplicationTasksRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationTasksFromDLQ provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetReplicationTasksFromDLQ(ctx context.Context, request *persistence.GetReplicationTasksFromDLQRequest) (*persistence.GetReplicationTasksResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetReplicationTasksResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetReplicationTasksFromDLQRequest) *persistence.GetReplicationTasksResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetReplicationTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetReplicationTasksFromDLQRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShardID provides a mock function with given fields:
func (_m *ExecutionManager) GetShardID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetTimerIndexTasks provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetTimerIndexTasks(ctx context.Context, request *persistence.GetTimerIndexTasksRequest) (*persistence.GetTimerIndexTasksResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetTimerIndexTasksResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetTimerIndexTasksRequest) *persistence.GetTimerIndexTasksResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetTimerIndexTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetTimerIndexTasksRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransferTasks provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetTransferTasks(ctx context.Context, request *persistence.GetTransferTasksRequest) (*persistence.GetTransferTasksResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetTransferTasksResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetTransferTasksRequest) *persistence.GetTransferTasksResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetTransferTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetTransferTasksRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) GetWorkflowExecution(ctx context.Context, request *persistence.GetWorkflowExecutionRequest) (*persistence.GetWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.GetWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.GetWorkflowExecutionRequest) *persistence.GetWorkflowExecutionResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.GetWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsWorkflowExecutionExists provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) IsWorkflowExecutionExists(ctx context.Context, request *persistence.IsWorkflowExecutionExistsRequest) (*persistence.IsWorkflowExecutionExistsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.IsWorkflowExecutionExistsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.IsWorkflowExecutionExistsRequest) *persistence.IsWorkflowExecutionExistsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.IsWorkflowExecutionExistsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.IsWorkflowExecutionExistsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListConcreteExecutions provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) ListConcreteExecutions(ctx context.Context, request *persistence.ListConcreteExecutionsRequest) (*persistence.ListConcreteExecutionsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.ListConcreteExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.ListConcreteExecutionsRequest) *persistence.ListConcreteExecutionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListConcreteExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.ListConcreteExecutionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCurrentExecutions provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) ListCurrentExecutions(ctx context.Context, request *persistence.ListCurrentExecutionsRequest) (*persistence.ListCurrentExecutionsResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.ListCurrentExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.ListCurrentExecutionsRequest) *persistence.ListCurrentExecutionsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.ListCurrentExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.ListCurrentExecutionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutReplicationTaskToDLQ provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) PutReplicationTaskToDLQ(ctx context.Context, request *persistence.PutReplicationTaskToDLQRequest) error {
	ret := _m.Called(ctx, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.PutReplicationTaskToDLQRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RangeCompleteCrossClusterTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) RangeCompleteCrossClusterTask(ctx context.Context, request *persistence.RangeCompleteCrossClusterTaskRequest) (*persistence.RangeCompleteCrossClusterTaskResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.RangeCompleteCrossClusterTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.RangeCompleteCrossClusterTaskRequest) *persistence.RangeCompleteCrossClusterTaskResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.RangeCompleteCrossClusterTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.RangeCompleteCrossClusterTaskRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RangeCompleteReplicationTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) RangeCompleteReplicationTask(ctx context.Context, request *persistence.RangeCompleteReplicationTaskRequest) (*persistence.RangeCompleteReplicationTaskResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.RangeCompleteReplicationTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.RangeCompleteReplicationTaskRequest) *persistence.RangeCompleteReplicationTaskResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.RangeCompleteReplicationTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.RangeCompleteReplicationTaskRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RangeCompleteTimerTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) RangeCompleteTimerTask(ctx context.Context, request *persistence.RangeCompleteTimerTaskRequest) (*persistence.RangeCompleteTimerTaskResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.RangeCompleteTimerTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.RangeCompleteTimerTaskRequest) *persistence.RangeCompleteTimerTaskResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.RangeCompleteTimerTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.RangeCompleteTimerTaskRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RangeCompleteTransferTask provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) RangeCompleteTransferTask(ctx context.Context, request *persistence.RangeCompleteTransferTaskRequest) (*persistence.RangeCompleteTransferTaskResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.RangeCompleteTransferTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.RangeCompleteTransferTaskRequest) *persistence.RangeCompleteTransferTaskResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.RangeCompleteTransferTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.RangeCompleteTransferTaskRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RangeDeleteReplicationTaskFromDLQ provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) RangeDeleteReplicationTaskFromDLQ(ctx context.Context, request *persistence.RangeDeleteReplicationTaskFromDLQRequest) (*persistence.RangeDeleteReplicationTaskFromDLQResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.RangeDeleteReplicationTaskFromDLQResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.RangeDeleteReplicationTaskFromDLQRequest) *persistence.RangeDeleteReplicationTaskFromDLQResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.RangeDeleteReplicationTaskFromDLQResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.RangeDeleteReplicationTaskFromDLQRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkflowExecution provides a mock function with given fields: ctx, request
func (_m *ExecutionManager) UpdateWorkflowExecution(ctx context.Context, request *persistence.UpdateWorkflowExecutionRequest) (*persistence.UpdateWorkflowExecutionResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *persistence.UpdateWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *persistence.UpdateWorkflowExecutionRequest) *persistence.UpdateWorkflowExecutionResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.UpdateWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *persistence.UpdateWorkflowExecutionRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
