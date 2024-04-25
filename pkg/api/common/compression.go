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

const (
	CompressionTypeNone = 0
	CompressionTypeLZMA = 100
)

var CompressionTypeEnum = NewEnum()

func init() {
	CompressionTypeEnum.MustRegister("none", CompressionTypeNone)
	CompressionTypeEnum.MustRegister("lzma", CompressionTypeLZMA)
}

var (
	CompressionNone *Compression = nil
	CompressionLZMA *Compression = NewCompression(CompressionTypeLZMA)
)

// NewCompression
func NewCompression(_type int32) *Compression {
	return &Compression{
		Type: _type,
	}
}

// Ensure returns new or existing Compression
func (x *Compression) Ensure(_type int32) *Compression {
	if x == nil {
		return NewCompression(_type)
	}
	return x
}

// GetName
func (x *Compression) GetName() string {
	return CompressionTypeEnum.GetName(x.GetType())
}

// String
func (x *Compression) String() string {
	return "no be implemented"
}
