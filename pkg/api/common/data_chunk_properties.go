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

// NewDataChunkProperties
func NewDataChunkProperties() *DataChunkProperties {
	return &DataChunkProperties{}
}

// HasDigest
func (x *DataChunkProperties) HasDigest() bool {
	if x == nil {
		return false
	}
	return x.Digest != nil
}

// SetDigest
func (x *DataChunkProperties) SetDigest(digest *Digest) *DataChunkProperties {
	if x == nil {
		return nil
	}
	x.Digest = digest
	return x
}

// HasOffset
func (x *DataChunkProperties) HasOffset() bool {
	if x == nil {
		return false
	}
	return x.Offset != nil
}

// SetOffset
func (x *DataChunkProperties) SetOffset(offset int64) *DataChunkProperties {
	if x == nil {
		return nil
	}
	x.Offset = new(int64)
	*x.Offset = offset
	return x
}

// HasLen
func (x *DataChunkProperties) HasLen() bool {
	if x == nil {
		return false
	}
	return x.Len != nil
}

// SetLen
func (x *DataChunkProperties) SetLen(len int64) *DataChunkProperties {
	if x == nil {
		return nil
	}
	x.Len = new(int64)
	*x.Len = len
	return x
}

// HasTotal
func (x *DataChunkProperties) HasTotal() bool {
	if x == nil {
		return false
	}
	return x.Total != nil
}

// SetTotal
func (x *DataChunkProperties) SetTotal(total int64) *DataChunkProperties {
	if x == nil {
		return nil
	}
	x.Total = new(int64)
	*x.Total = total
	return x
}

// HasLast
func (x *DataChunkProperties) HasLast() bool {
	if x == nil {
		return false
	}
	return x.Last != nil
}

// SetLast
func (x *DataChunkProperties) SetLast(last bool) *DataChunkProperties {
	if x == nil {
		return nil
	}
	x.Last = new(bool)
	*x.Last = last
	return x
}

// String
func (x *DataChunkProperties) String() string {
	return "no be implemented"
}
