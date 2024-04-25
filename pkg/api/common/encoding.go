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

// NewEncoding
func NewEncoding(method ...string) *Encoding {
	f := new(Encoding)
	if len(method) > 0 {
		f.Set(method[0])
	}
	return f
}

// Set
func (x *Encoding) Set(method string) *Encoding {
	if x == nil {
		return nil
	}
	x.Method = method
	return x
}

// Equals
func (x *Encoding) Equals(encoding *Encoding) bool {
	if x == nil {
		return false
	}
	if encoding == nil {
		return false
	}
	return x.GetMethod() == encoding.GetMethod()
}

// String
func (x *Encoding) String() string {
	if x == nil {
		return ""
	}
	return x.GetMethod()
}
