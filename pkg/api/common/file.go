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

import "fmt"

// NewFile
func NewFile() *File {
	return new(File)
}

// Len
func (x *File) Len() int {
	if x != nil {
		return len(x.Data)
	}
	return 0
}

// SetFilename
func (x *File) SetFilename(filename string) *File {
	if x == nil {
		return nil
	}
	x.Filename = x.Filename.Ensure().Set(filename)
	return x
}

// SetData
func (x *File) SetData(data []byte) *File {
	if x == nil {
		return nil
	}
	x.Data = data
	return x
}

// String
func (x *File) String() string {
	if x == nil {
		return ""
	}
	return fmt.Sprintf("%s[%d]", x.Filename.String(), x.Len())
}
