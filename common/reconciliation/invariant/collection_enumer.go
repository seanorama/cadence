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

// Code generated by "enumer -type=Collection"; DO NOT EDIT.

package invariant

import (
	"fmt"
	"strings"
)

const _CollectionName = "CollectionMutableStateCollectionHistoryCollectionDomain"

var _CollectionIndex = [...]uint8{0, 22, 39, 55}

const _CollectionLowerName = "collectionmutablestatecollectionhistorycollectiondomain"

func (i Collection) String() string {
	if i < 0 || i >= Collection(len(_CollectionIndex)-1) {
		return fmt.Sprintf("Collection(%d)", i)
	}
	return _CollectionName[_CollectionIndex[i]:_CollectionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CollectionNoOp() {
	var x [1]struct{}
	_ = x[CollectionMutableState-(0)]
	_ = x[CollectionHistory-(1)]
	_ = x[CollectionDomain-(2)]
}

var _CollectionValues = []Collection{CollectionMutableState, CollectionHistory, CollectionDomain}

var _CollectionNameToValueMap = map[string]Collection{
	_CollectionName[0:22]:       CollectionMutableState,
	_CollectionLowerName[0:22]:  CollectionMutableState,
	_CollectionName[22:39]:      CollectionHistory,
	_CollectionLowerName[22:39]: CollectionHistory,
	_CollectionName[39:55]:      CollectionDomain,
	_CollectionLowerName[39:55]: CollectionDomain,
}

var _CollectionNames = []string{
	_CollectionName[0:22],
	_CollectionName[22:39],
	_CollectionName[39:55],
}

// CollectionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CollectionString(s string) (Collection, error) {
	if val, ok := _CollectionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CollectionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Collection values", s)
}

// CollectionValues returns all values of the enum
func CollectionValues() []Collection {
	return _CollectionValues
}

// CollectionStrings returns a slice of all String values of the enum
func CollectionStrings() []string {
	strs := make([]string, len(_CollectionNames))
	copy(strs, _CollectionNames)
	return strs
}

// IsACollection returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Collection) IsACollection() bool {
	for _, v := range _CollectionValues {
		if i == v {
			return true
		}
	}
	return false
}
