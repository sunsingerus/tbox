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

package sections

import (
	"fmt"
	"github.com/sunsingerus/tbox/pkg/config/items"
)

// OTPConfigurator - one time password (token) related section of config
type OTPConfigurator interface {
	GetOTP() items.OTPConfigurator
}

// Interface compatibility
var _ OTPConfigurator = OTP{}

// OTP
type OTP struct {
	OTP *items.OTP `mapstructure:"otp"`
}

// OTPNormalize is a normalizer
func (o OTP) OTPNormalize() OTP {
	if o.OTP == nil {
		o.OTP = items.NewOTP()
	}
	return o
}

// GetOTP is a getter
func (o OTP) GetOTP() items.OTPConfigurator {
	return o.OTP
}

// String is a stringifier
func (o OTP) String() string {
	return fmt.Sprintf("OTP=%s", o.OTP)
}
