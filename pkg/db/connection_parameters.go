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

package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mailru/go-clickhouse"

	"github.com/sunsingerus/tbox/pkg/config/sections"
)

var (
	DefaultConnectTimeout        = 10 * time.Second
	DefaultQueryTimeout          = 10 * time.Second
	DefaultExecTimeout           = 10 * time.Second
	DefaultMaxOpenConnections    = 100
	DefaultMaxIdleConnections    = 1
	DefaultConnectionMaxLifetime = 10 * time.Minute
	DefaultConnectionMaxIdleTime = 1 * time.Second
)

// ConnectionParameters specifies database connection parameters
type ConnectionParameters struct {
	// driverName which to connect with
	DriverName string

	// username which to connect with
	Username string
	// password which to connect with
	Password string

	// hostname where to connect to
	Hostname string
	// port where to connect to
	Port int

	// database to work with
	Database string

	// Ready-to-use DSN string
	Dsn string
	// DSN string with hidden credentials. Can be used in logs, etc
	DsnHiddenCredentials string

	// ConnectTimeout specifies connect timeout [OPTIONAL]
	ConnectTimeout *time.Duration
	// QueryTimeout specifies query timeout [OPTIONAL]
	QueryTimeout *time.Duration
	// ExecTimeout specifies exec timeout [OPTIONAL]
	ExecTimeout *time.Duration

	// MaxOpenConnections specifies max number of open connections [OPTIONAL]
	MaxOpenConnections *int
	// MaxIdleConnections specifies max number of idling connections [OPTIONAL]
	MaxIdleConnections *int
	// ConnectionMaxLifeTime specifies max connection lifetime [OPTIONAL]
	ConnectionMaxLifetime *time.Duration
	// ConnectionMaxIdleTime specifies max connection idling time [OPTIONAL]
	ConnectionMaxIdleTime *time.Duration
}

// NewConnectionParameters creates new database conenction parameters
func NewConnectionParameters(driverName, username, password string, hostname string, port int, database, dsn string) *ConnectionParameters {
	if dsn == "" {
		// Construct DSN from components
		return &ConnectionParameters{
			DriverName:           driverName,
			Username:             username,
			Password:             password,
			Hostname:             hostname,
			Port:                 port,
			Database:             database,
			Dsn:                  makeDSN(driverName, username, password, hostname, port, database, false),
			DsnHiddenCredentials: makeDSN(driverName, username, password, hostname, port, database, true),
		}
	} else {
		// Have DSN provided
		return &ConnectionParameters{
			DriverName:           driverName,
			Dsn:                  dsn,
			DsnHiddenCredentials: hideCredentialsDSN(driverName, dsn),
		}
	}
}

// NewConnectionParametersConfig creates new database connection parameters from config
func NewConnectionParametersConfig(driverName string, untyped interface{}) *ConnectionParameters {
	switch driverName {
	case "mysql":
		cfg := untyped.(sections.MySQLConfigurator)
		return NewConnectionParameters(
			driverName,
			cfg.GetMySQLUsername(),
			cfg.GetMySQLPassword(),
			cfg.GetMySQLHostname(),
			cfg.GetMySQLPort(),
			cfg.GetMySQLDatabase(),
			cfg.GetMySQLDSN(),
		)
	case "pgx":
		cfg := untyped.(sections.PostgreSQLConfigurator)
		return NewConnectionParameters(
			driverName,
			cfg.GetPostgreSQLUsername(),
			cfg.GetPostgreSQLPassword(),
			cfg.GetPostgreSQLHostname(),
			cfg.GetPostgreSQLPort(),
			cfg.GetPostgreSQLDatabase(),
			cfg.GetPostgreSQLDSN(),
		)
	case "clickhouse":
		cfg := untyped.(sections.ClickHouseConfigurator)
		return NewConnectionParameters(
			driverName,
			cfg.GetClickHouseUsername(),
			cfg.GetClickHousePassword(),
			cfg.GetClickHouseHostname(),
			cfg.GetClickHousePort(),
			cfg.GetClickHouseDatabase(),
			cfg.GetClickHouseDSN(),
		)
	}
	return nil

}

// GetDriverName gets name of the driver
func (c *ConnectionParameters) GetDriverName() string {
	if c == nil {
		return ""
	}
	return c.DriverName
}

// SetDriverName sets name of the driver
func (c *ConnectionParameters) SetDriverName(name string) {
	if c == nil {
		return
	}
	c.DriverName = name
}

// GetDSN gets DSN
func (c *ConnectionParameters) GetDSN() string {
	if c == nil {
		return ""
	}
	return c.Dsn
}

// GetDSNWithHiddenCredentials gets DSN with hidden credentials. Handy for logging
func (c *ConnectionParameters) GetDSNWithHiddenCredentials() string {
	if c == nil {
		return ""
	}
	return c.DsnHiddenCredentials
}

func (c *ConnectionParameters) HasConnectTimeout() bool {
	if c == nil {
		return false
	}
	return c.ConnectTimeout != nil
}

func (c *ConnectionParameters) GetConnectTimeout() time.Duration {
	if c.HasConnectTimeout() {
		return *c.ConnectTimeout
	}
	return DefaultConnectTimeout
}

func (c *ConnectionParameters) SetConnectTimeout(t time.Duration) {
	if c == nil {
		return
	}
	c.ConnectTimeout = &t
}

func (c *ConnectionParameters) HasQueryTimeout() bool {
	if c == nil {
		return false
	}
	return c.QueryTimeout != nil
}

func (c *ConnectionParameters) GetQueryTimeout() time.Duration {
	if c.HasQueryTimeout() {
		return *c.QueryTimeout
	}
	return DefaultQueryTimeout
}

func (c *ConnectionParameters) SetQueryTimeout(t time.Duration) {
	if c == nil {
		return
	}
	c.QueryTimeout = &t
}

func (c *ConnectionParameters) HasExecTimeout() bool {
	if c == nil {
		return false
	}
	return c.ExecTimeout != nil
}

func (c *ConnectionParameters) GetExecTimeout() time.Duration {
	if c.HasExecTimeout() {
		return *c.ExecTimeout
	}
	return DefaultExecTimeout
}

func (c *ConnectionParameters) SetExecTimeout(t time.Duration) {
	if c == nil {
		return
	}
	c.ExecTimeout = &t
}

func (c *ConnectionParameters) HasMaxOpenConnections() bool {
	if c == nil {
		return false
	}
	return c.MaxOpenConnections != nil
}

func (c *ConnectionParameters) GetMaxOpenConnections() int {
	if c.HasMaxOpenConnections() {
		return *c.MaxOpenConnections
	}
	return DefaultMaxOpenConnections
}

func (c *ConnectionParameters) SetMaxOpenConnections(m int) {
	if c == nil {
		return
	}
	c.MaxOpenConnections = &m
}

func (c *ConnectionParameters) HasMaxIdleConnections() bool {
	if c == nil {
		return false
	}
	return c.MaxIdleConnections != nil
}

func (c *ConnectionParameters) GetMaxIdleConnections() int {
	if c.HasMaxIdleConnections() {
		return *c.MaxIdleConnections
	}
	return DefaultMaxIdleConnections
}

func (c *ConnectionParameters) SetMaxIdleConnections(m int) {
	if c == nil {
		return
	}
	c.MaxIdleConnections = &m
}

func (c *ConnectionParameters) HasConnectionMaxLifetime() bool {
	if c == nil {
		return false
	}
	return c.ConnectionMaxLifetime != nil
}

func (c *ConnectionParameters) GetConnectionMaxLifetime() time.Duration {
	if c.HasConnectionMaxLifetime() {
		return *c.ConnectionMaxLifetime
	}
	return DefaultConnectionMaxLifetime
}

func (c *ConnectionParameters) SetConnectionMaxLifetime(m time.Duration) {
	if c == nil {
		return
	}
	c.ConnectionMaxLifetime = &m
}

func (c *ConnectionParameters) HasConnectionMaxIdleTime() bool {
	if c == nil {
		return false
	}
	return c.ConnectionMaxIdleTime != nil
}

func (c *ConnectionParameters) GetConnectionMaxIdleTime() time.Duration {
	if c.HasConnectionMaxIdleTime() {
		return *c.ConnectionMaxIdleTime
	}
	return DefaultConnectionMaxIdleTime
}

func (c *ConnectionParameters) SetConnectionMaxIdleTime(m time.Duration) {
	if c == nil {
		return
	}
	c.ConnectionMaxIdleTime = &m
}
