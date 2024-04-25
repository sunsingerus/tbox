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

//
//
// Wrap AddressMap
//
//

// Has wraps AddressMap.Has
func (x *Metadata) Has(domain ...*Domain) bool {
	return x.GetAddresses().Has(domain...)
}

// Set wraps AddressMap.Set
func (x *Metadata) Set(entities ...interface{}) *Metadata {
	x.EnsureAddresses().Set(entities...)
	return x
}

// Append wraps AddressMap.Append
func (x *Metadata) Append(entities ...interface{}) *Metadata {
	x.EnsureAddresses().Append(entities...)
	return x
}

// Address customization
//
// SetUuidFromString
func (x *Metadata) SetUuidFromString(id string, domain ...*Domain) *Metadata {
	var domain0 = DomainThis
	var domain1 = DomainUUID
	switch len(domain) {
	case 2:
		domain0 = domain[0]
		domain1 = domain[1]
	case 1:
		domain0 = domain[0]
	}
	x.Set(domain0, domain1, NewAddressUuidFromString(id, domain1))
	return x
}

// SetRandomUuid domains are optional
func (x *Metadata) SetRandomUuid(domains ...*Domain) *Metadata {
	// Default values
	var domain0 = DomainThis
	var domain1 = DomainUUID

	switch len(domains) {
	case 2:
		domain0 = domains[0]
		domain1 = domains[1]
	case 1:
		domain0 = domains[0]
	}
	x.Set(domain0, domain1, NewAddressUuidRandom(domain1))
	return x
}

//
//
// Wrap Address
//
//

// SetS3
func (x *Metadata) SetS3(bucket, object string) *Metadata {
	return x.Set(DomainThis, DomainS3, NewAddress().Set(NewS3Address(bucket, object)))
}

// SetKafka
func (x *Metadata) SetKafka(topic string, partition int32) *Metadata {
	return x.Set(DomainThis, DomainKafka, NewAddress().Set(NewKafkaAddress(topic, partition)))
}

// SetUUID
func (x *Metadata) SetUUID(uuid *UUID) *Metadata {
	return x.Set(DomainThis, DomainUUID, NewAddress().Set(uuid))
}

// GetUUID
func (x *Metadata) GetUuid() *UUID {
	return x.GetAddresses().First(DomainThis, DomainUUID).GetUuid()
}

// SetUserID
func (x *Metadata) SetUserID(userID *UserID) *Metadata {
	return x.Set(DomainThis, DomainUserID, NewAddress().Set(userID))
}

// GetUserID
func (x *Metadata) GetUserID() *UserID {
	return x.GetAddresses().First(DomainThis, DomainUserID).GetUserId()
}

// SetDirname
func (x *Metadata) SetDirname(dirname string) *Metadata {
	return x.Set(DomainThis, DomainDirname, NewAddress().Set(NewDirname(dirname)))
}

// GetDirname
func (x *Metadata) GetDirname() string {
	return x.GetAddresses().First(DomainThis, DomainDirname).GetDirname().String()
}

// SetFilename
func (x *Metadata) SetFilename(filename string) *Metadata {
	return x.Set(DomainThis, DomainFilename, NewAddress().Set(NewFilename(filename)))
}

// GetFilename
func (x *Metadata) GetFilename() string {
	return x.GetAddresses().First(DomainThis, DomainFilename).GetFilename().String()
}

// SetURL
func (x *Metadata) SetURL(url string) *Metadata {
	return x.Set(DomainThis, DomainURL, NewAddress().Set(NewURL(url)))
}

// SetDomain
func (x *Metadata) SetDomain(domain *Domain) *Metadata {
	return x.Set(DomainThis, DomainDomain, NewAddress().Set(domain))
}

// GetDomain
func (x *Metadata) GetDomain() *Domain {
	return x.GetAddresses().First(DomainThis, DomainDomain).GetDomain()
}

// SetMachineID
func (x *Metadata) SetMachineID(machineID *MachineID) *Metadata {
	return x.Set(DomainThis, DomainMachineID, NewAddress().Set(machineID))
}

// GetMachineID
func (x *Metadata) GetMachineID() *MachineID {
	return x.GetAddresses().First(DomainThis, DomainMachineID).GetMachineId()
}

// SetAssetID
func (x *Metadata) SetAssetID(assetID *MachineID) *Metadata {
	return x.Set(DomainThis, DomainAssetID, NewAddress().Set(assetID))
}

// GetAssetID
func (x *Metadata) GetAssetID() *MachineID {
	return x.GetAddresses().First(DomainThis, DomainAssetID).GetMachineId()
}

// SetCustom
func (x *Metadata) SetCustom(s string) *Metadata {
	return x.Set(DomainThis, DomainCustom, NewAddress().Set(s))
}

// GetCustom
func (x *Metadata) GetCustom() string {
	return x.GetAddresses().First(DomainThis, DomainCustom).GetCustom()
}

// GetContextUUID
func (x *Metadata) GetContextUuid() *UUID {
	return x.GetAddresses().First(DomainContext, DomainUUID).GetUuid()
}

// SetContextUUID
func (x *Metadata) SetContextUUID(uuid *UUID) *Metadata {
	return x.Set(DomainContext, DomainUUID, NewAddress().Set(uuid))
}

// GetTaskUUID
func (x *Metadata) GetTaskUuid() *UUID {
	return x.GetAddresses().First(DomainTask, DomainUUID).GetUuid()
}

// SetTaskUUID
func (x *Metadata) SetTaskUUID(uuid *UUID) *Metadata {
	return x.Set(DomainTask, DomainUUID, NewAddress().Set(uuid))
}

// GetResultDomain
func (x *Metadata) GetResultDomain() *Domain {
	return x.GetAddresses().First(DomainResult, DomainDomain).GetDomain()
}

// SetResultDomain
func (x *Metadata) SetResultDomain(domain *Domain) *Metadata {
	return x.Set(DomainResult, DomainDomain, NewAddress().Set(domain))
}

// GetAddress
func (x *Metadata) GetAddress() *UUID {
	return x.GetAddresses().First(DomainAddress, DomainAddress).GetUuid()
}

// SetAddress
func (x *Metadata) SetAddress(address *Address) *Metadata {
	return x.Set(DomainAddress, DomainAddress, address)
}
