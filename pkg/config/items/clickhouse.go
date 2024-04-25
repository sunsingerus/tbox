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
type ClickHouse struct {
	// Enabled specifies whether item is enabled
	Enabled bool `mapstructure:"enabled"`
	// Schema specifies schema to use Ex.: https http tcp
	Schema string `mapstructure:"schema"`
	// Username specifies ClickHouse username
	Username string `mapstructure:"username"`
	// Password specifies ClickHouse password
	Password string `mapstructure:"password"`
	// Hostname specifies ClickHouse host
	Hostname string `mapstructure:"hostname"`
	// Port specifies ClickHouse port
	Port int `mapstructure:"port"`
	// Database specifies ClickHouse database
	Database string `mapstructure:"database"`
	// DSN in the form: http://username:password@clickhouse.host:8123/database as a combination of all above
	DSN string `mapstructure:"dsn"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewClickHouse
func NewClickHouse() *ClickHouse {
	return new(ClickHouse)
}

// GetEnabled
func (c *ClickHouse) GetEnabled() bool {
	if c == nil {
		return false
	}
	return c.Enabled
}

// GetSchema
func (c *ClickHouse) GetSchema() string {
	if c == nil {
		return ""
	}
	return c.Schema
}

// GetUsername
func (c *ClickHouse) GetUsername() string {
	if c == nil {
		return ""
	}
	return c.Username
}

// GetPassword
func (c *ClickHouse) GetPassword() string {
	if c == nil {
		return ""
	}
	return c.Password
}

// GetHostname
func (c *ClickHouse) GetHostname() string {
	if c == nil {
		return ""
	}
	return c.Hostname
}

// GetPort
func (c *ClickHouse) GetPort() int {
	if c == nil {
		return 0
	}
	return c.Port
}

// GetDatabase
func (c *ClickHouse) GetDatabase() string {
	if c == nil {
		return ""
	}
	return c.Database
}

// GetDSN
func (c *ClickHouse) GetDSN() string {
	if c == nil {
		return ""
	}
	return c.DSN
}

// String
func (c *ClickHouse) String() string {
	if c == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", c.Enabled)
	_, _ = fmt.Fprintf(b, "Username: %v\n", c.Username)
	_, _ = fmt.Fprintf(b, "Password: %v\n", c.Password)
	_, _ = fmt.Fprintf(b, "Hostname: %v\n", c.Hostname)
	_, _ = fmt.Fprintf(b, "Port: %v\n", c.Port)
	_, _ = fmt.Fprintf(b, "DSN: %v\n", c.DSN)

	return b.String()
}
