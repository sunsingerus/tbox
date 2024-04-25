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

// TLSConfigurator
type TLSConfigurator interface {
	IsEnabled() bool
	GetEnabled() bool
	HasServerNameOverride() bool
	GetServerNameOverride() string
	HasCAFile() bool
	GetCAFile() string
	HasPrivateKeyFile() bool
	GetPrivateKeyFile() string
	HasPublicCertFile() bool
	GetPublicCertFile() string
}

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type TLS struct {
	Enabled            bool   `mapstructure:"enabled"`
	ServerNameOverride string `mapstructure:"server-name-override"`
	CAFile             string `mapstructure:"ca-file"`
	PrivateKeyFile     string `mapstructure:"private-key-file"`
	PublicCertFile     string `mapstructure:"public-cert-file"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

var _ TLSConfigurator = &TLS{}

// NewTLS creates new TLS
func NewTLS() *TLS {
	return new(TLS)
}

// IsEnabled checks whether TLS is enabled
func (t *TLS) IsEnabled() bool {
	return t.GetEnabled()
}

// GetEnabled is a getter
func (t *TLS) GetEnabled() bool {
	if t == nil {
		return false
	}
	return t.Enabled
}

// HasServerNameOverride checks whether ServerNameOverride is specified
func (t *TLS) HasServerNameOverride() bool {
	return t.GetServerNameOverride() != ""
}

// GetServerNameOverride is a getter
func (t *TLS) GetServerNameOverride() string {
	if t == nil {
		return ""
	}
	return t.ServerNameOverride
}

// HasCAFile checks whether CAFile is specified
func (t *TLS) HasCAFile() bool {
	return t.GetCAFile() != ""
}

// GetCAFile is a getter
func (t *TLS) GetCAFile() string {
	if t == nil {
		return ""
	}
	return t.CAFile
}

// HasPrivateKeyFile checks whether PrivateKeyFile is specified
func (t *TLS) HasPrivateKeyFile() bool {
	return t.GetPrivateKeyFile() != ""
}

// GetPrivateKeyFile is a getter
func (t *TLS) GetPrivateKeyFile() string {
	if t == nil {
		return ""
	}
	return t.PrivateKeyFile
}

// HasPublicCertFile checks whether PublicCertFile is specified
func (t *TLS) HasPublicCertFile() bool {
	return t.GetPublicCertFile() != ""
}

// GetPublicCertFile is a getter
func (t *TLS) GetPublicCertFile() string {
	if t == nil {
		return ""
	}
	return t.PublicCertFile
}

// String is a stringifier
func (t *TLS) String() string {
	if t == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", t.Enabled)
	_, _ = fmt.Fprintf(b, "CAFile: %v\n", t.CAFile)
	_, _ = fmt.Fprintf(b, "ServerNameOverride: %v\n", t.ServerNameOverride)
	_, _ = fmt.Fprintf(b, "PrivateKeyFile: %v\n", t.PrivateKeyFile)
	_, _ = fmt.Fprintf(b, "PublicCertFile: %v\n", t.PublicCertFile)

	return b.String()
}
