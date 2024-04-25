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

import (
	"fmt"
)

// Ensure interface compatibility
var (
	_ IDataChunk = &DataChunk{}
)

func NewDataChunk() *DataChunk {
	return new(DataChunk)
}

//
// Interface functions
//

// GetDataLen is an IDataChunk interface function
func (x *DataChunk) GetDataLen() int {
	if x == nil {
		return 0
	}
	return len(x.GetData())
}

// SetData is an IDataChunk interface function
func (x *DataChunk) SetData(data []byte) {
	if x == nil {
		return
	}
	x.Data = data
}

// GetLast is an IDataChunk interface function
func (x *DataChunk) GetLast() bool {
	return x.GetProperties().GetLast()
}

// SetLast is an DataChunk interface function
func (x *DataChunk) SetLast(last bool) {
	x.EnsureProperties().SetLast(last)
}

// GetOffset is an DataChunk interface function
func (x *DataChunk) GetOffset() int64 {
	return x.GetProperties().GetOffset()
}

// SetOffset is an DataChunk interface function
func (x *DataChunk) SetOffset(offset int64) {
	x.EnsureProperties().SetOffset(offset)
}

// EnsureProperties is a getter with guaranteed result
func (x *DataChunk) EnsureProperties() *DataChunkProperties {
	if x.GetProperties() == nil {
		x.SetProperties(NewDataChunkProperties())
	}
	return x.GetProperties()
}

// SetProperties is a setter
func (x *DataChunk) SetProperties(props *DataChunkProperties) {
	if x == nil {
		return
	}
	x.Properties = props
}

// String is a stringifier
func (x *DataChunk) String() string {
	// Fetch offset of this chunk within the stream
	offset := "not specified"
	if x.GetProperties().HasOffset() {
		offset = fmt.Sprintf("%d", x.GetProperties().GetOffset())
	}

	last := "not specified"
	if x.GetProperties().HasLast() {
		last = fmt.Sprintf("%v", x.GetProperties().GetLast())
	}

	return fmt.Sprintf("len:%d, offset:%s, last:%v",
		x.GetDataLen(),
		offset,
		last,
	)
}
