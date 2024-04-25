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

package items

import (
	"encoding/json"
)

// OTPConfigurator - one time password (token) related section of config
type OTPConfigurator interface {
	GetEnabled() bool
	GetServer() string
	GetClient() string
}

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type OTP struct {
	Enabled bool   `mapstructure:"enabled"`
	Server  string `mapstructure:"server"`
	Client  string `mapstructure:"client"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

var _ OTPConfigurator = &OTP{}

// NewOTP is a condtructor
func NewOTP() *OTP {
	return &OTP{}
}

// GetEnabled is a getter
func (o *OTP) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

// GetServer is a getter
func (o *OTP) GetServer() string {
	if o == nil {
		return ""
	}
	return o.Server
}

// GetClient is a getter
func (o *OTP) GetClient() string {
	if o == nil {
		return ""
	}
	return o.Client
}

// String is a stringifier
func (o *OTP) String() string {
	self, _ := json.Marshal(o)
	return string(self)
}
