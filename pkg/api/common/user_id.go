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

import "github.com/google/uuid"

// NewUserID
func NewUserID() *UserID {
	return &UserID{}
}

// NewRandomUserID
func NewRandomUserID() *UserID {
	return NewUserID().SetString(uuid.New().String())
}

// NewUserIDFromString
func NewUserIDFromString(str string) *UserID {
	return NewUserID().SetString(str)
}

// SetBytes
func (x *UserID) SetBytes(bytes []byte) *UserID {
	x.Data = bytes
	return x
}

// GetBytes
func (x *UserID) GetBytes() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// SetString
func (x *UserID) SetString(str string) *UserID {
	if x == nil {
		return nil
	}
	x.Data = []byte(str)
	return x
}

// String
func (x *UserID) String() string {
	if x != nil {
		return string(x.Data)
	}
	return ""
}
