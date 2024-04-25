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

package mysql

import (
	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/db"
)

const driverName = "mysql"

// NewConnection creates new MySQL connection
func NewConnection(params *db.ConnectionParameters) *db.Connection {
	// Do not perform connection immediately, do it in lazy manner
	params.SetDriverName(driverName)
	return db.NewConnection(params)
}

// NewConnectionParametersConfig creates new MySQL connection parameters from config
func NewConnectionParametersConfig(cfg sections.MySQLConfigurator) *db.ConnectionParameters {
	return db.NewConnectionParametersConfig(driverName, cfg)
}

// NewConnectionConfig creates new MySQL connection from config
func NewConnectionConfig(cfg sections.MySQLConfigurator) *db.Connection {
	return db.NewConnection(NewConnectionParametersConfig(cfg))
}
