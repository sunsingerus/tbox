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

// NewObjectRequest is a constructor
func NewObjectRequest() *ObjectRequest {
	return new(ObjectRequest)
}

// SetRequestDomain is a setter
func (x *ObjectRequest) SetRequestDomain(domain *Domain) *ObjectRequest {
	if x == nil {
		return nil
	}
	x.RequestDomain = domain
	return x
}

// SetResultDomain is a setter
func (x *ObjectRequest) SetResultDomain(domain *Domain) *ObjectRequest {
	if x == nil {
		return nil
	}
	x.ResultDomain = domain
	return x
}

// GetAddress is a convenience wrapper to get one address of a specified domain
func (x *ObjectRequest) GetAddress(domains ...*Domain) *Address {
	if x == nil {
		return nil
	}
	return x.GetAddresses().First(domains...)
}

// SetAddress is a convenience wrapper to set one address with optional domain(s)
// The format is: [domain,] address, [address, ...]
func (x *ObjectRequest) SetAddress(entities ...interface{}) *ObjectRequest {
	if x == nil {
		return nil
	}
	x.Addresses = x.Addresses.Ensure()
	x.Addresses.Set(entities...)
	return x
}

// AppendAddress is a convenience wrapper to append oen address with optional domain(s)
// The format is: [domain,] address, [address, ...]
func (x *ObjectRequest) AppendAddress(entities ...interface{}) *ObjectRequest {
	if x == nil {
		return nil
	}
	x.Addresses = x.Addresses.Ensure()
	x.Addresses.Append(entities...)
	return x
}

// String is a stringifier
func (x *ObjectRequest) String() string {
	return "to be implemented"
}
