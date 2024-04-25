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

// PostgreSQLConfigurator
type PostgreSQLConfigurator interface {
	GetPostgreSQLUsername() string
	GetPostgreSQLPassword() string
	GetPostgreSQLHostname() string
	GetPostgreSQLPort() int
	GetPostgreSQLDatabase() string
	GetPostgreSQLDSN() string
}

// Interface compatibility
var _ PostgreSQLConfigurator = PostgreSQL{}

// PostgreSQL
type PostgreSQL struct {
	PostgreSQL *items.PostgreSQL `mapstructure:"postgresql"`
}

// PostgreSQLNormalize
func (c PostgreSQL) PostgreSQLNormalize() PostgreSQL {
	if c.PostgreSQL == nil {
		c.PostgreSQL = items.NewPostgreSQL()
	}
	return c
}

// GetPostgreSQLUsername
func (c PostgreSQL) GetPostgreSQLUsername() string {
	return c.PostgreSQL.GetUsername()
}

// GetPostgreSQLPassword
func (c PostgreSQL) GetPostgreSQLPassword() string {
	return c.PostgreSQL.GetPassword()
}

// GetPostgreSQLHostname
func (c PostgreSQL) GetPostgreSQLHostname() string {
	return c.PostgreSQL.GetHostname()
}

// GetPostgreSQLPort
func (c PostgreSQL) GetPostgreSQLPort() int {
	return c.PostgreSQL.GetPort()
}

// GetPostgreSQLDatabase
func (c PostgreSQL) GetPostgreSQLDatabase() string {
	return c.PostgreSQL.GetDatabase()
}

// GetPostgreSQLDSN
func (c PostgreSQL) GetPostgreSQLDSN() string {
	return c.PostgreSQL.GetDSN()
}

// String
func (c PostgreSQL) String() string {
	return fmt.Sprintf("PostgreSQL=%s", c.PostgreSQL)
}
