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
// Wrap metadata
//

// GetType gets task type
func (x *Task) GetType() int32 {
	return x.GetHeader().GetType()
}

// SetType
func (x *Task) SetType(_type int32) *Task {
	x.EnsureHeader().SetType(_type)
	return x
}

// GetName gets task name
func (x *Task) GetName() string {
	return x.GetHeader().GetName()
}

// SetName
func (x *Task) SetName(name string) *Task {
	x.EnsureHeader().SetName(name)
	return x
}

// GetStatus gets task status
func (x *Task) GetStatus() int32 {
	return x.GetHeader().GetStatus()
}

// SetStatus sets task status
func (x *Task) SetStatus(status int32) *Task {
	x.EnsureHeader().SetStatus(status)
	return x
}

// GetUuid
func (x *Task) GetUuid() *UUID {
	return x.GetHeader().GetAddresses().First(DomainThis, DomainUUID).GetUuid()
}

// GetUuidAsString
func (x *Task) GetUuidAsString() string {
	return x.GetUuid().String()
}

// SetUuid
func (x *Task) SetUuid(address *Address) *Task {
	x.EnsureHeader().Set(DomainThis, DomainUUID, address)
	return x
}

// SetUuidFromString
func (x *Task) SetUuidFromString(id string) *Task {
	x.SetUuid(NewAddressUuidFromString(id, DomainUUID))
	return x
}

// SetUserID
func (x *Task) SetUserID(address *Address) *Task {
	x.EnsureHeader().Set(DomainUser, DomainUserID, address)
	return x
}

// SetUserIDFromString
func (x *Task) SetUserIDFromString(id string) *Task {
	x.SetUserID(NewAddressUserIDFromString(id, DomainUserID))
	return x
}

// GetUserID
func (x *Task) GetUserID() *UserID {
	return x.GetHeader().GetAddresses().First(DomainUser, DomainUserID).GetUserId()
}

// GetUserIDAsString
func (x *Task) GetUserIDAsString() string {
	return x.GetUserID().String()
}

// SetEmail
func (x *Task) SetEmail(address *Address) *Task {
	x.EnsureHeader().Set(DomainUser, DomainEmail, address)
	return x
}

// SetEmailFromString
func (x *Task) SetEmailFromString(email string) *Task {
	x.SetEmail(NewAddressEmailFromString(email, DomainEmail))
	return x
}

// GetEmail
func (x *Task) GetEmail() *Email {
	return x.GetHeader().GetAddresses().First(DomainUser, DomainEmail).GetEmail()
}

// GetEmailAsString
func (x *Task) GetEmailAsString() string {
	return x.GetEmail().String()
}

// CreateUuid creates new random UUID
func (x *Task) CreateUuid() *Task {
	return x.SetUuid(NewAddressUuidRandom(DomainUUID))
}

// GetReferenceUuid
func (x *Task) GetReferenceUuid() *UUID {
	return x.GetHeader().GetAddresses().First(DomainReference, DomainUUID).GetUuid()
}

// GetReferenceUuidAsString
func (x *Task) GetReferenceUuidAsString() string {
	return x.GetReferenceUuid().String()
}

// SetReferenceUuid
func (x *Task) SetReferenceUuid(uuid *UUID) *Task {
	x.EnsureHeader().EnsureAddresses().Set(DomainReference, DomainUUID, NewAddress(uuid))
	return x
}

// SetReferenceUuidFromString
func (x *Task) SetReferenceUuidFromString(id string) *Task {
	x.SetReferenceUuid(NewUuidFromString(id))
	return x
}

// GetContextUuid
func (x *Task) GetContextUuid() *UUID {
	return x.GetHeader().GetAddresses().First(DomainContext, DomainUUID).GetUuid()
}

// GetContextUuidAsString
func (x *Task) GetContextUuidAsString() string {
	return x.GetContextUuid().String()
}

// SetContextUuid
func (x *Task) SetContextUuid(uuid *UUID) *Task {
	x.EnsureHeader().EnsureAddresses().Set(DomainContext, DomainUUID, NewAddress(uuid))
	return x
}

// SetContextUuidFromString
func (x *Task) SetContextUuidFromString(id string) *Task {
	x.SetContextUuid(NewUuidFromString(id))
	return x
}

// GetResult
func (x *Task) GetResult() *Address {
	return x.GetHeader().GetAddresses().First(DomainResult)
}

// GetResults
func (x *Task) GetResults() []*Address {
	return x.GetHeader().GetAddresses().All(DomainResult)
}

// AppendResult
func (x *Task) AppendResult(address *Address) *Task {
	x.EnsureHeader().EnsureAddresses().Append(DomainResult, address)
	return x
}

// SetResult
func (x *Task) SetResult(address *Address) *Task {
	x.EnsureHeader().EnsureAddresses().Set(DomainResult, address.Domain(), address)
	return x
}

// GetDescription
func (x *Task) GetDescription() string {
	return x.GetHeader().GetDescription()
}

// SetDescription
func (x *Task) SetDescription(description string) *Task {
	x.EnsureHeader().SetDescription(description)
	return x
}
