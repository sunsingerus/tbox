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

// ClickHouseConfigurator
type ClickHouseConfigurator interface {
	GetClickHouseSchema() string
	GetClickHouseUsername() string
	GetClickHousePassword() string
	GetClickHouseHostname() string
	GetClickHousePort() int
	GetClickHouseDatabase() string
	GetClickHouseDSN() string
}

// Interface compatibility
var _ ClickHouseConfigurator = ClickHouse{}

// ClickHouse
type ClickHouse struct {
	ClickHouse *items.ClickHouse `mapstructure:"clickhouse"`
}

// ClickHouseNormalize
func (c ClickHouse) ClickHouseNormalize() ClickHouse {
	if c.ClickHouse == nil {
		c.ClickHouse = items.NewClickHouse()
	}
	return c
}

// GetClickHouseSchema
func (c ClickHouse) GetClickHouseSchema() string {
	return c.ClickHouse.GetSchema()
}

// GetClickHouseUsername
func (c ClickHouse) GetClickHouseUsername() string {
	return c.ClickHouse.GetUsername()
}

// GetClickHousePassword
func (c ClickHouse) GetClickHousePassword() string {
	return c.ClickHouse.GetPassword()
}

// GetClickHouseHostname
func (c ClickHouse) GetClickHouseHostname() string {
	return c.ClickHouse.GetHostname()
}

// GetClickHousePort
func (c ClickHouse) GetClickHousePort() int {
	return c.ClickHouse.GetPort()
}

// GetClickHouseDatabase
func (c ClickHouse) GetClickHouseDatabase() string {
	return c.ClickHouse.GetDatabase()
}

// GetClickHouseDSN
func (c ClickHouse) GetClickHouseDSN() string {
	return c.ClickHouse.GetDSN()
}

// String
func (c ClickHouse) String() string {
	return fmt.Sprintf("ClickHouse=%s", c.ClickHouse)
}
