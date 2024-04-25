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

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type Manticore struct {
	Enabled bool `mapstructure:"enabled"`
	// Hostname specifies ClickHouse host
	Hostname string `mapstructure:"hostname"`
	// Port specifies ClickHouse port
	Port int `mapstructure:"port"`
	// DSN in the form: manticore.host:port as a combination of all above
	DSN string `mapstructure:"dsn"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewManticore is a constructor
func NewManticore() *Manticore {
	return new(Manticore)
}

// GetEnabled is a getter
func (c *Manticore) GetEnabled() bool {
	if c == nil {
		return false
	}
	return c.Enabled
}

// GetHostname is a getter
func (c *Manticore) GetHostname() string {
	if c == nil {
		return ""
	}
	return c.Hostname
}

// GetPort is a getter
func (c *Manticore) GetPort() int {
	if c == nil {
		return 0
	}
	return c.Port
}

// GetDSN is a getter
func (c *Manticore) GetDSN() string {
	if c == nil {
		return ""
	}
	return c.DSN
}

// String is a stringifier
func (c *Manticore) String() string {
	if c == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", c.Enabled)
	_, _ = fmt.Fprintf(b, "Hostname: %v\n", c.Hostname)
	_, _ = fmt.Fprintf(b, "Port: %v\n", c.Port)
	_, _ = fmt.Fprintf(b, "DSN: %v\n", c.DSN)

	return b.String()
}
