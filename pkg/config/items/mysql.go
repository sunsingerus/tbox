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
type MySQL struct {
	Enabled bool `mapstructure:"enabled"`
	// Username specifies MySQL username
	Username string `mapstructure:"username"`
	// Password specifies MySQL password
	Password string `mapstructure:"password"`
	// Hostname specifies MySQL host
	Hostname string `mapstructure:"hostname"`
	// Port specifies MySQL port
	Port int `mapstructure:"port"`
	// Database specifies MySQL database
	Database string `mapstructure:"database"`
	// DSN in the form: user:password@tcp(host:3306)/database as a combination of all above
	DSN string `mapstructure:"dsn"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewMySQL is a constructor
func NewMySQL() *MySQL {
	return new(MySQL)
}

// GetEnabled is a getter
func (c *MySQL) GetEnabled() bool {
	if c == nil {
		return false
	}
	return c.Enabled
}

// GetUsername is a getter
func (c *MySQL) GetUsername() string {
	if c == nil {
		return ""
	}
	return c.Username
}

// GetPassword is a getter
func (c *MySQL) GetPassword() string {
	if c == nil {
		return ""
	}
	return c.Password
}

// GetHostname is a getter
func (c *MySQL) GetHostname() string {
	if c == nil {
		return ""
	}
	return c.Hostname
}

// GetPort is a getter
func (c *MySQL) GetPort() int {
	if c == nil {
		return 0
	}
	return c.Port
}

// GetDatabase is a getter
func (c *MySQL) GetDatabase() string {
	if c == nil {
		return ""
	}
	return c.Database
}

// GetDSN is a getter
func (c *MySQL) GetDSN() string {
	if c == nil {
		return ""
	}
	return c.DSN
}

// String is a stringifier
func (c *MySQL) String() string {
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
