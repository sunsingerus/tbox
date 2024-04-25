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
	"bytes"
	"fmt"
)

type ServiceConfigurator interface {
	GetEnabled() bool
	GetName() string
	GetType() ServiceType
	GetAddress() string
	GetPaths() PathsConfigurator
	GetTLS() TLSConfigurator
	GetOAuth() OAuthConfigurator
	GetOTP() OTPConfigurator
	GetBackend() ServiceConfigurator
	fmt.Stringer
}

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type Service struct {
	Enabled bool        `mapstructure:"enabled"`
	Name    string      `mapstructure:"name"`
	Type    ServiceType `mapstructure:"type"`
	Address string      `mapstructure:"address"`
	Paths   *MultiPaths `mapstructure:"paths"`
	TLS     *TLS        `mapstructure:"tls"`
	OAuth   *OAuth      `mapstructure:"oauth"`
	OTP     *OTP        `mapstructure:"otp"`
	Backend *Service    `mapstructure:"backend"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

var _ ServiceConfigurator = &Service{}

// NewService creates new Service
func NewService() *Service {
	return new(Service)
}

// GetEnabled is a getter
func (s *Service) GetEnabled() bool {
	if s == nil {
		return false
	}
	return s.Enabled
}

// GetName is a getter
func (s *Service) GetName() string {
	if s == nil {
		return ""
	}
	return s.Name
}

// GetType is a getter
func (s *Service) GetType() ServiceType {
	if s == nil {
		return ""
	}
	return s.Type
}

// GetAddress is a getter
func (s *Service) GetAddress() string {
	if s == nil {
		return ""
	}
	return s.Address
}

// GetPaths is a getter
func (s *Service) GetPaths() PathsConfigurator {
	if s == nil {
		return nil
	}
	return s.Paths
}

// GetTLS is a getter
func (s *Service) GetTLS() TLSConfigurator {
	if s == nil {
		return nil
	}
	return s.TLS
}

// GetOAuth is a getter
func (s *Service) GetOAuth() OAuthConfigurator {
	if s == nil {
		return nil
	}
	return s.OAuth
}

// GetOTP is a getter
func (s *Service) GetOTP() OTPConfigurator {
	if s == nil {
		return nil
	}
	return s.OTP
}

// GetBackend is a getter
func (s *Service) GetBackend() ServiceConfigurator {
	if s == nil {
		return nil
	}
	return s.Backend
}

// String is a stringifier
func (s *Service) String() string {
	if s == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Name: %s\nType: %s\nAddress: %s\nPaths: %s\nTLS: %s\nOAuth: %s\nBackend %s\n", s.Name, s.Type, s.Address, s.Paths, s.TLS, s.OAuth, s.Backend)

	return b.String()
}
