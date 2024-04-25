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

// MySQLConfigurator
type MySQLConfigurator interface {
	GetMySQLUsername() string
	GetMySQLPassword() string
	GetMySQLHostname() string
	GetMySQLPort() int
	GetMySQLDatabase() string
	GetMySQLDSN() string
}

// Interface compatibility
var _ MySQLConfigurator = MySQL{}

// MySQL
type MySQL struct {
	MySQL *items.MySQL `mapstructure:"mysql"`
}

// MySQLNormalize
func (c MySQL) MySQLNormalize() MySQL {
	if c.MySQL == nil {
		c.MySQL = items.NewMySQL()
	}
	return c
}

// GetMySQLUsername
func (c MySQL) GetMySQLUsername() string {
	return c.MySQL.GetUsername()
}

// GetMySQLPassword
func (c MySQL) GetMySQLPassword() string {
	return c.MySQL.GetPassword()
}

// GetMySQLHostname
func (c MySQL) GetMySQLHostname() string {
	return c.MySQL.GetHostname()
}

// GetMySQLPort
func (c MySQL) GetMySQLPort() int {
	return c.MySQL.GetPort()
}

// GetMySQLDatabase
func (c MySQL) GetMySQLDatabase() string {
	return c.MySQL.GetDatabase()
}

// GetMySQLDSN
func (c MySQL) GetMySQLDSN() string {
	return c.MySQL.GetDSN()
}

// String
func (c MySQL) String() string {
	return fmt.Sprintf("MySQL=%s", c.MySQL)
}
