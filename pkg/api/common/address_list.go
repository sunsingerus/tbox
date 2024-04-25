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

import "strings"

// NewAddressList creates new AddressList
func NewAddressList() *AddressList {
	return &AddressList{}
}

// NewAddressListFromString creates new AddressList from string
func NewAddressListFromString(str string) *AddressList {
	parts := strings.Split(str, separatorAddressList)
	if len(parts) == 0 {
		return nil
	}

	list := NewAddressList()
	for _, address := range parts {
		if address := NewAddressFromString(address); address != nil {
			list.Append(address)
		}
	}

	if list.Len() > 0 {
		return list
	}

	return nil
}

// Ensure returns new or existing AddressList
func (x *AddressList) Ensure() *AddressList {
	if x == nil {
		return NewAddressList()
	}
	return x
}

// Len returns either how many addresses of specified Domain are in the list
// or len of the whole AddressList in case no domain specified
func (x *AddressList) Len(domain ...*Domain) int {
	if len(domain) > 0 {
		return x.LenOf(domain[0])
	}
	return len(x.GetAddresses())
}

// LenOf returns how many addresses of specified Domain are in the list
func (x *AddressList) LenOf(domain *Domain) int {
	res := 0
	for _, address := range x.GetAddresses() {
		if address.Domain().Equals(domain) {
			res++
		}
	}
	return res
}

// All wraps GetAddresses and Select and returns all Addresses in specified domains
// May return nil
func (x *AddressList) All(domains ...*Domain) []*Address {
	if len(domains) > 0 {
		return x.Select(domains...).GetAddresses()
	}
	return x.GetAddresses()
}

// Slice returns slice of addresses
// May return nil
func (x *AddressList) Slice(a, b int) []*Address {
	// Sanity check 1
	if (a < 0) || (b < 0) {
		return nil
	}
	if a > b {
		return nil
	}

	// Sanity check 2
	addresses := x.GetAddresses()
	if (a > len(addresses)) || (b > len(addresses)) {
		return nil
	}

	// Boundaries look like sane, continue with addresses

	if len(addresses) == 0 {
		return nil
	}

	return addresses[a:b]
}

// Has checks whether AddressList has something (in case do domain specified) or specified nested domain exists
func (x *AddressList) Has(domain ...*Domain) bool {
	if len(domain) > 0 {
		return x.HasDomain(domain[0])
	}
	// Deal with AddressList itself
	return x.Len() > 0
}

// First returns the first address in the list of specified nested domain
// May return nil
func (x *AddressList) First(domain ...*Domain) *Address {
	if len(domain) > 0 {
		return x.FirstOf(domain[0])
	}
	// Deal with AddressList itself
	if addresses := x.GetAddresses(); len(addresses) > 0 {
		return addresses[0]
	}
	return nil
}

// Last returns the last address in the list of specified nested domain
func (x *AddressList) Last(domain ...*Domain) *Address {
	if len(domain) > 0 {
		return x.LastOf(domain[0])
	}
	// Deal with AddressList itself
	if addresses := x.GetAddresses(); len(addresses) > 0 {
		return addresses[len(addresses)-1]
	}
	return nil
}

// HasDomain checks whether specified domain exists
func (x *AddressList) HasDomain(domain *Domain) bool {
	return x.LenOf(domain) > 0
}

// FirstOf gets the first address of specified domains
func (x *AddressList) FirstOf(domains ...*Domain) *Address {
	for _, address := range x.GetAddresses() {
		for _, domain := range domains {
			if address.Domain().Equals(domain) {
				return address
			}
		}
	}
	return nil
}

// LastOf gets the last address of specified domains
func (x *AddressList) LastOf(domains ...*Domain) *Address {
	var res *Address = nil
	for _, address := range x.GetAddresses() {
		for _, domain := range domains {
			if address.Domain().Equals(domain) {
				res = address
			}
		}
	}
	return res
}

// Select selects all addresses with specified domains into new AddressList.
// May return nil in case nothing is selected.
func (x *AddressList) Select(domains ...*Domain) *AddressList {
	var res *AddressList = nil
	for _, address := range x.GetAddresses() {
		for _, domain := range domains {
			if address.Domain().Equals(domain) {
				if res == nil {
					res = NewAddressList()
				}
				res.Append(address)
			}
		}
	}
	return res
}

// Exclude selects all addresses without specified domains into new AddressList
// May return nil in case result is empty
func (x *AddressList) Exclude(domains ...*Domain) *AddressList {
	var res *AddressList = nil
	for _, address := range x.GetAddresses() {
		for _, domain := range domains {
			if address.Domain().Equals(domain) {
				// Skip this address, it is in delete list
			} else {
				// Keep this address
				if res == nil {
					res = NewAddressList()
				}
				res.Append(address)
			}
		}
	}
	return res
}

// Delete deletes from the AddressList all addresses with specified domains
func (x *AddressList) Delete(domains ...*Domain) *AddressList {
	var keep []*Address = nil
	for _, address := range x.GetAddresses() {
		for _, domain := range domains {
			if address.Domain().Equals(domain) {
				// Skip this address, it is in delete list
			} else {
				// Keep this address
				keep = append(keep, address)
			}
		}
	}
	x.Replace(keep...)
	return x
}

// Append appends addresses to the AddressList
func (x *AddressList) Append(addresses ...*Address) *AddressList {
	if x != nil {
		x.Addresses = append(x.Addresses, addresses...)
		return x
	}
	return nil
}

// Replace replaces existing list with provided list of addresses
func (x *AddressList) Replace(addresses ...*Address) *AddressList {
	if x != nil {
		x.Addresses = append([]*Address{}, addresses...)
		return x
	}
	return nil
}

const separatorAddressList = ","

// String stringifies address list writing addresses w/o domains
func (x *AddressList) String() string {
	res := ""
	for _, address := range x.GetAddresses() {
		if len(res) > 0 {
			res += separatorAddressList
		}
		res += address.String()
	}
	return res
}

// FullString stringifies address list writing addresses with domains
func (x *AddressList) FullString() string {
	res := ""
	for _, address := range x.GetAddresses() {
		if len(res) > 0 {
			res += separatorAddressList
		}
		res += address.FullString()
	}
	return res
}
