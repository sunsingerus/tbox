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

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"time"
)

// NewMetadata
func NewMetadata() *Metadata {
	return new(Metadata)
}

// HasType
func (x *Metadata) HasType() bool {
	if x == nil {
		return false
	}
	return x.Type != nil
}

// SetType
func (x *Metadata) SetType(_type int32) *Metadata {
	if x == nil {
		return nil
	}
	x.Type = new(int32)
	*x.Type = _type
	return x
}

// HasName
func (x *Metadata) HasName() bool {
	if x == nil {
		return false
	}
	return x.Name != nil
}

// SetName
func (x *Metadata) SetName(name string) *Metadata {
	if x == nil {
		return nil
	}
	x.Name = new(string)
	*x.Name = name
	return x
}

// HasVersion
func (x *Metadata) HasVersion() bool {
	if x == nil {
		return false
	}
	return x.Version != nil
}

// SetVersion
func (x *Metadata) SetVersion(version int32) *Metadata {
	if x == nil {
		return nil
	}
	x.Version = new(int32)
	*x.Version = version
	return x
}

// HasDescription
func (x *Metadata) HasDescription() bool {
	if x == nil {
		return false
	}
	return x.Description != nil
}

// SetDescription
func (x *Metadata) SetDescription(description string) *Metadata {
	if x == nil {
		return nil
	}
	x.Description = new(string)
	*x.Description = description
	return x
}

// HasStatus
func (x *Metadata) HasStatus() bool {
	if x == nil {
		return false
	}
	return x.Status != nil
}

// SetStatus
func (x *Metadata) SetStatus(status int32) *Metadata {
	if x == nil {
		return nil
	}
	x.Status = new(int32)
	*x.Status = status
	return x
}

// HasMode
func (x *Metadata) HasMode() bool {
	if x == nil {
		return false
	}
	return x.Mode != nil
}

// SetMode
func (x *Metadata) SetMode(mode int32) *Metadata {
	if x == nil {
		return nil
	}
	x.Mode = new(int32)
	*x.Mode = mode
	return x
}

// HasTimestamp
func (x *Metadata) HasTimestamp() bool {
	if x == nil {
		return false
	}
	return x.Ts != nil
}

// SetTimestamp
func (x *Metadata) SetTimestamp(seconds int64, nanos int32) *Metadata {
	if x == nil {
		return nil
	}
	x.Ts = new(timestamp.Timestamp)
	x.Ts.Seconds = seconds
	x.Ts.Nanos = nanos
	return x
}

// CreateTimestamp creates current timestamp
func (x *Metadata) CreateTimestamp() *Metadata {
	now := time.Now()
	seconds := now.Unix()           // seconds since 1970
	nanoseconds := now.Nanosecond() // nanosecond offset within the second
	return x.SetTimestamp(seconds, int32(nanoseconds))
}

// HasAddresses
func (x *Metadata) HasAddresses() bool {
	if x == nil {
		return false
	}
	return x.Addresses != nil
}

// SetAddresses
func (x *Metadata) SetAddresses(addresses *AddressMap) *Metadata {
	if x == nil {
		return nil
	}
	x.Addresses = addresses
	return x
}

// EnsureAddresses
func (x *Metadata) EnsureAddresses() *AddressMap {
	if x == nil {
		return nil
	}
	if x.HasAddresses() {
		return x.GetAddresses()
	}
	x.SetAddresses(NewAddressMap())
	return x.GetAddresses()
}

// HasProperties
func (x *Metadata) HasProperties() bool {
	if x == nil {
		return false
	}
	return x.Properties != nil
}

// SetProperties
func (x *Metadata) SetProperties(properties *DataChunkProperties) *Metadata {
	if x == nil {
		return nil
	}
	x.Properties = properties
	return x
}

// EnsureProperties
func (x *Metadata) EnsureProperties() *DataChunkProperties {
	if x == nil {
		return nil
	}
	if x.HasProperties() {
		return x.GetProperties()
	}
	x.SetProperties(NewDataChunkProperties())
	return x.GetProperties()
}

// HasTypes
func (x *Metadata) HasTypes() bool {
	if x == nil {
		return false
	}
	return x.Types != nil
}

// SetTypes
func (x *Metadata) SetTypes(types *SliceInt32) *Metadata {
	if x == nil {
		return nil
	}
	x.Types = types
	return x
}

// EnsureTypes
func (x *Metadata) EnsureTypes() *SliceInt32 {
	if x == nil {
		return nil
	}
	if x.HasTypes() {
		return x.GetTypes()
	}
	x.SetTypes(NewSliceInt32())
	return x.GetTypes()
}

// HasVersions
func (x *Metadata) HasVersions() bool {
	if x == nil {
		return false
	}
	return x.Versions != nil
}

// SetVersions
func (x *Metadata) SetVersions(versions *SliceInt32) *Metadata {
	if x == nil {
		return nil
	}
	x.Versions = versions
	return x
}

// EnsureVersions
func (x *Metadata) EnsureVersions() *SliceInt32 {
	if x == nil {
		return nil
	}
	if x.HasVersions() {
		return x.GetVersions()
	}
	x.SetVersions(NewSliceInt32())
	return x.GetVersions()
}

// HasStatuses
func (x *Metadata) HasStatuses() bool {
	if x == nil {
		return false
	}
	return x.Statuses != nil
}

// SetStatuses
func (x *Metadata) SetStatuses(statuses *SliceInt32) *Metadata {
	if x == nil {
		return nil
	}
	x.Statuses = statuses
	return x
}

// EnsureStatuses
func (x *Metadata) EnsureStatuses() *SliceInt32 {
	if x == nil {
		return nil
	}
	if x.HasStatuses() {
		return x.GetStatuses()
	}
	x.SetStatuses(NewSliceInt32())
	return x.GetStatuses()
}

// HasModes
func (x *Metadata) HasModes() bool {
	if x == nil {
		return false
	}
	return x.Modes != nil
}

// SetModes
func (x *Metadata) SetModes(modes *SliceInt32) *Metadata {
	if x == nil {
		return nil
	}
	x.Modes = modes
	return x
}

// EnsureModes
func (x *Metadata) EnsureModes() *SliceInt32 {
	if x == nil {
		return nil
	}
	if x.HasModes() {
		return x.GetModes()
	}
	x.SetModes(NewSliceInt32())
	return x.GetModes()
}

// Log
func (x *Metadata) Log() {
	if x == nil {
		return
	}
	log.Infof("metadata: %s", x.String())
}

// String
func (x *Metadata) String() string {
	if yml, err := yaml.Marshal(x); err == nil {
		return string(yml)
	} else {
		return "unable to stringify metadata"
	}
}
