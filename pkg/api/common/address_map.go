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

// NewAddressMap creates new AddressMap
func NewAddressMap() *AddressMap {
	return new(AddressMap)
}

// Ensure returns new or existing AddressMap
func (x *AddressMap) Ensure() *AddressMap {
	if x == nil {
		return NewAddressMap()
	}
	return x
}

// GetList gets specified AddressList of specified domain
func (x *AddressMap) GetList(domain *Domain) *AddressList {
	if mp := x.GetMap(); mp != nil {
		if list, ok := mp[domain.GetName()]; ok {
			return list
		}
	}
	return nil
}

// GetLists gets all AddressList from the AddressMap
func (x *AddressMap) GetLists() []*AddressList {
	if mp := x.GetMap(); mp != nil {
		var res []*AddressList
		for _, list := range mp {
			if res == nil {
				res = make([]*AddressList, x.Len())
			}
			res = append(res, list)
		}
		return res
	}
	return nil
}

// Collapse collapses all lists from this map into one list
func (x *AddressMap) Collapse() *AddressList {
	lists := NewAddressList()
	for _, list := range x.GetLists() {
		lists.Append(list.GetAddresses()...)
	}
	if lists.Len() > 0 {
		return lists
	}
	return nil
}

// EnsureList makes sure AddressList of specified domain exists.
// It uses already existing domain AddressList or creates new if none found
func (x *AddressMap) EnsureList(domain *Domain) *AddressList {
	if x == nil {
		return nil
	}
	if x.Has(domain) {
		return x.GetList(domain)
	}
	return x.NewList(domain)
}

// NewList creates new AddressList of specified domain. Existing one will be overwritten.
func (x *AddressMap) NewList(domain *Domain) *AddressList {
	if x == nil {
		return nil
	}
	return x.SetList(domain, NewAddressList()).GetList(domain)
}

// SetList sets AddressList of specified domain. Existing one will be overwritten.
func (x *AddressMap) SetList(domain *Domain, list *AddressList) *AddressMap {
	if x == nil {
		return nil
	}
	x.ensureMap()
	x.Map[domain.GetName()] = list
	return x
}

// ensureMap makes sure map is created
func (x *AddressMap) ensureMap() {
	if x == nil {
		// Unable to ensure map inside nil struct
		return
	}
	if x.GetMap() == nil {
		// Ensure map exists
		x.Map = make(map[string]*AddressList)
	}
}

// String stringifies AddressMap
func (x *AddressMap) String() string {
	return "no be implemented"
}
