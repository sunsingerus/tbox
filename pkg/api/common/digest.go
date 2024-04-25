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

import "encoding/hex"

// NewDigest creates bew digest
func NewDigest() *Digest {
	return new(Digest)
}

// SetType sets digest type
func (x *Digest) SetType(_type DigestType) *Digest {
	if x == nil {
		return nil
	}
	x.Type = _type
	return x
}

// SetData sets data directly
func (x *Digest) SetData(data []byte) *Digest {
	if x == nil {
		return nil
	}
	x.Data = data
	return x
}

// SetDataFromString sets data from string value
func (x *Digest) SetDataFromString(data string) *Digest {
	if x == nil {
		return nil
	}
	if bin, err := hex.DecodeString(data); err == nil {
		x.Data = bin
	} else {
		x.Data = nil
	}
	return x
}

// String
func (x *Digest) String() string {
	if x == nil {
		return ""
	}
	if len(x.Data) == 0 {
		return ""
	}
	return hex.EncodeToString(x.Data)
}
