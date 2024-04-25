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

// NewFilename
func NewFilename(filename ...string) *Filename {
	f := new(Filename)
	if len(filename) > 0 {
		f.Set(filename[0])
	}
	return f
}

// Ensure
func (x *Filename) Ensure() *Filename {
	if x != nil {
		return x
	}
	return new(Filename)
}

// Set
func (x *Filename) Set(filename string) *Filename {
	if x == nil {
		return nil
	}
	x.Filename = filename
	return x
}

// Equals
func (x *Filename) Equals(filename *Filename) bool {
	if x == nil {
		return false
	}
	if filename == nil {
		return false
	}
	return x.GetFilename() == filename.GetFilename()
}

// String
func (x *Filename) String() string {
	if x == nil {
		return ""
	}
	return x.GetFilename()
}
