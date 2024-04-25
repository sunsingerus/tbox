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

// NewDomain creates new Domain
func NewDomain(name ...string) *Domain {
	d := new(Domain)
	if len(name) > 0 {
		d.SetName(name[0])
	}
	return d
}

// Ensure returns new or existing Domain
func (x *Domain) Ensure() *Domain {
	if x == nil {
		return NewDomain()
	}
	return x
}

// SetName sets Domain name
func (x *Domain) SetName(name string) *Domain {
	if x == nil {
		return nil
	}
	x.Name = name
	return x
}

// Equals checks whether Domains are equal internally
func (x *Domain) Equals(domain *Domain) bool {
	if x == nil {
		return false
	}
	if domain == nil {
		return false
	}
	return x.GetName() == domain.GetName()
}

// String returns string representation of a Domain
func (x *Domain) String() string {
	if x == nil {
		return ""
	}
	return x.GetName()
}
