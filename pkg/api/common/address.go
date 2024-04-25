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

// AddressType represents all types of domain-specific addresses in the system
const (
	// AddressReserved specifies reserved value due to first enum value has to be zero in proto3
	AddressReserved int32 = 0
	// AddressS3 specifies S3 and MinIO address
	AddressS3 int32 = 100
	// AddressKafka specifies Kafka address
	AddressKafka int32 = 200
	// AddressDigest specifies digest-based address
	AddressDigest int32 = 300
	// AddressUUID specifies UUID-based address
	AddressUUID int32 = 400
	// AddressUserID specifies UserID-based address. Used to specify any related user (owner, sender, etc)
	AddressUserID int32 = 500
	// AddressDirname specifies dirname/path-based address
	AddressDirname int32 = 600
	// AddressFilename specifies filename/filepath-based address
	AddressFilename int32 = 700
	// AddressURL specifies URL address
	AddressURL int32 = 800
	// AddressDomain specifies Domain address
	AddressDomain int32 = 900
	// AddressMachineID specifies MachineID-based address
	AddressMachineID int32 = 1000
	// AddressEmail specifies Email address
	AddressEmail int32 = 1100
	// AddressCustom specifies Custom string
	AddressCustom int32 = 1200
)

var AddressTypeEnum = NewEnum()

func init() {
	AddressTypeEnum.MustRegister("AddressReserved", AddressReserved)
	AddressTypeEnum.MustRegister("AddressS3", AddressS3)
	AddressTypeEnum.MustRegister("AddressKafka", AddressKafka)
	AddressTypeEnum.MustRegister("AddressDigest", AddressDigest)
	AddressTypeEnum.MustRegister("AddressUUID", AddressUUID)
	AddressTypeEnum.MustRegister("AddressUserID", AddressUserID)
	AddressTypeEnum.MustRegister("AddressDirname", AddressDirname)
	AddressTypeEnum.MustRegister("AddressFilename", AddressFilename)
	AddressTypeEnum.MustRegister("AddressURL", AddressURL)
	AddressTypeEnum.MustRegister("AddressDomain", AddressDomain)
	AddressTypeEnum.MustRegister("AddressMachineID", AddressMachineID)
	AddressTypeEnum.MustRegister("AddressEmail", AddressEmail)
	AddressTypeEnum.MustRegister("AddressCustom", AddressCustom)
}

// NewAddress creates new Address with specified domain
// Call example:
// NewAddress(domain, address)
func NewAddress(entities ...interface{}) *Address {
	var domain *Domain = nil
	var address interface{} = nil
	for _, entity := range entities {
		switch typed := entity.(type) {
		case *Domain:
			domain = typed
		default:
			address = entity
		}
	}

	res := new(Address)
	if domain != nil {
		res.SetExplicitDomain(domain)
	}
	if address != nil {
		res.Set(address)
	}
	return res
}

// NewAddressUuidRandom creates new Address with specified Domain with random UUID
func NewAddressUuidRandom(domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewUuidRandom())
}

// NewAddressUuidFromString creates new Address with specified Domain with UUID fetched from string
func NewAddressUuidFromString(str string, domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewUuid().SetString(str))
}

// NewAddressUserIDFromString creates new Address with specified Domain with UUID fetched from string
func NewAddressUserIDFromString(str string, domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewUserID().SetString(str))
}

// NewAddressEmailFromString creates new Address with specified Domain with Email fetched from string
func NewAddressEmailFromString(str string, domain ...interface{}) *Address {
	return NewAddress(domain...).Set(NewEmail().SetString(str))
}

// NewAddressFromString creates new address from string by parsing this string according to address's stringify rules.
func NewAddressFromString(str string) *Address {
	parts := strings.SplitN(str, separatorAddress, 2)
	if len(parts) != 2 {
		return nil
	}
	domain := DomainFromString(parts[0])
	if domain == nil {
		return nil
	}
	switch domain {
	case DomainS3:
		return NewAddress(NewS3AddressFromString(parts[1]))
	case DomainKafka:
		return nil
	case DomainDigest:
		return nil
	case DomainUUID:
		return NewAddress(NewUuidFromString(parts[1]))
	case DomainUserID:
		return nil
	case DomainDirname:
		return nil
	case DomainFilename:
		return nil
	case DomainURL:
		return nil
	case DomainDomain:
		return nil
	case DomainMachineID:
		return nil
	case DomainEmail:
		return nil
	case DomainCustom:
		return nil
	}
	return nil
}

// Ensure returns new or existing Address
func (x *Address) Ensure() *Address {
	if x == nil {
		return NewAddress()
	}
	return x
}

// SetExplicitDomain is a setter
func (x *Address) SetExplicitDomain(domain *Domain) *Address {
	if x == nil {
		return nil
	}
	x.ExplicitDomain = domain
	return x
}

// Domain gets domain of an address - either explicitly specified or from the value
func (x *Address) Domain() *Domain {
	if x == nil {
		return nil
	}
	if explicit := x.GetExplicitDomain(); explicit != nil {
		return explicit
	}

	switch {
	case x.GetS3() != nil:
		return DomainS3
	case x.GetKafka() != nil:
		return DomainKafka
	case x.GetDigest() != nil:
		return DomainDigest
	case x.GetUuid() != nil:
		return DomainUUID
	case x.GetUserId() != nil:
		return DomainUserID
	case x.GetDirname() != nil:
		return DomainDirname
	case x.GetFilename() != nil:
		return DomainFilename
	case x.GetUrl() != nil:
		return DomainURL
	case x.GetDomain() != nil:
		return DomainDomain
	case x.GetMachineId() != nil:
		return DomainMachineID
	case x.GetEmail() != nil:
		return DomainEmail
	default:
		return DomainCustom
	}
}

// Set sets value of the Address
func (x *Address) Set(address interface{}) *Address {
	if x == nil {
		return nil
	}
	switch typed := address.(type) {
	case isAddress_Address:
		x.Address = typed
	case *S3Address:
		i := new(Address_S3)
		i.S3 = typed
		x.Address = i
	case *KafkaAddress:
		i := new(Address_Kafka)
		i.Kafka = typed
		x.Address = i
	case *Digest:
		i := new(Address_Digest)
		i.Digest = typed
		x.Address = i
	case *UUID:
		i := new(Address_Uuid)
		i.Uuid = typed
		x.Address = i
	case *UserID:
		i := new(Address_UserId)
		i.UserId = typed
		x.Address = i
	case *Dirname:
		i := new(Address_Dirname)
		i.Dirname = typed
		x.Address = i
	case *Filename:
		i := new(Address_Filename)
		i.Filename = typed
		x.Address = i
	case *URL:
		i := new(Address_Url)
		i.Url = typed
		x.Address = i
	case *Domain:
		i := new(Address_Domain)
		i.Domain = typed
		x.Address = i
	case *MachineID:
		i := new(Address_MachineId)
		i.MachineId = typed
		x.Address = i
	case *Email:
		i := new(Address_Email)
		i.Email = typed
		x.Address = i
	case string:
		i := new(Address_Custom)
		i.Custom = typed
		x.Address = i
	}
	return x
}

// String return string address w/o domain prefix
func (x *Address) String() string {
	if x == nil {
		return ""
	}

	if x.GetAddress() == nil {
		return "address unspecified"
	}

	switch {
	case x.GetS3() != nil:
		return x.GetS3().String()
	case x.GetKafka() != nil:
		return x.GetKafka().String()
	case x.GetDigest() != nil:
		return x.GetDigest().String()
	case x.GetUuid() != nil:
		return x.GetUuid().String()
	case x.GetUserId() != nil:
		return x.GetUserId().String()
	case x.GetDirname() != nil:
		return x.GetDirname().String()
	case x.GetFilename() != nil:
		return x.GetFilename().String()
	case x.GetUrl() != nil:
		return x.GetUrl().String()
	case x.GetDomain() != nil:
		return x.GetDomain().String()
	case x.GetMachineId() != nil:
		return x.GetMachineId().String()
	case x.GetEmail() != nil:
		return x.GetEmail().String()
	default:
		return x.GetCustom()
	}
}

const separatorAddress = "://"

// FullString returns string address with domain prefix
func (x *Address) FullString() string {
	domain := x.Domain()
	if domain == nil {
		return ""
	}
	return domain.GetName() + separatorAddress + x.String()
}
