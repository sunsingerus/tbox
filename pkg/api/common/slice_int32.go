// Copyright The TBox Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

// NewSliceInt32
func NewSliceInt32() *SliceInt32 {
	return new(SliceInt32)
}

// Ensure
func (x *SliceInt32) Ensure() *SliceInt32 {
	if x == nil {
		return NewSliceInt32()
	}
	return x
}

// Add
func (x *SliceInt32) Add(a ...int32) *SliceInt32 {
	x = x.Ensure()
	x.Slice = append(x.Slice, a...)
	return x
}

// Get
func (x *SliceInt32) Get(i int) int32 {
	if x == nil {
		return 0
	}
	if i >= len(x.Slice) {
		// Requested index is out of the scope
		return 0
	}

	return x.Slice[i]
}

// Fetch
func (x *SliceInt32) Fetch(i int) (int32, bool) {
	if x == nil {
		return 0, false
	}
	if i >= len(x.Slice) {
		// Requested index is out of the scope
		return 0, false
	}

	return x.Slice[i], true
}

// Len
func (x *SliceInt32) Len() int {
	if x == nil {
		return 0
	}
	return len(x.Slice)
}

// GetAll
func (x *SliceInt32) GetAll() []int32 {
	return x.GetSlice()
}

// String
func (x *SliceInt32) String() string {
	return "to be implemented"
}
