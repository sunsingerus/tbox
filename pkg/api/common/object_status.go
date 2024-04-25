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

// NewObjectStatus
func NewObjectStatus(status ...*Status) *ObjectStatus {
	d := new(ObjectStatus)
	if len(status) > 0 {
		d.SetStatus(status[0])
	}
	return d
}

// Ensure returns new or existing Status
func (x *ObjectStatus) Ensure() *ObjectStatus {
	if x == nil {
		return NewObjectStatus()
	}
	return x
}

// SetStatus sets status
func (x *ObjectStatus) SetStatus(status *Status) *ObjectStatus {
	if x == nil {
		return nil
	}
	x.Status = status
	return x
}

// SetDomain sets address
func (x *ObjectStatus) SetDomain(domain *Domain) *ObjectStatus {
	if x == nil {
		return nil
	}
	x.Domain = domain
	return x
}

// SetAddress sets address
func (x *ObjectStatus) SetAddress(address *Address) *ObjectStatus {
	if x == nil {
		return nil
	}
	x.Address = address
	return x
}

// String
func (x *ObjectStatus) String() string {
	return "to be implemented"
}
