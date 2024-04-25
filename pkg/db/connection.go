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
	"context"
	"fmt"
	"time"

	databasesql "github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"

	// register db drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mailru/go-clickhouse"
)

// Connection is a connection to a database
type Connection struct {
	params *ConnectionParameters
	conn   *databasesql.DB
}

// NewConnection creates new database connection
func NewConnection(params *ConnectionParameters) *Connection {
	// Do not perform connection immediately, do it in lazy manner
	return &Connection{
		params: params,
	}
}

// NewConnectionConfig creates new database connection from config
func NewConnectionConfig(driverName string, cfg interface{}) *Connection {
	return NewConnection(NewConnectionParametersConfig(driverName, cfg))
}

func (c *Connection) GetParams() *ConnectionParameters {
	if c == nil {
		return nil
	}
	return c.params
}

// Close closes open connection
func (c *Connection) Close() error {
	log.Debug("Close() Connection")
	if c == nil {
		return nil
	}
	if c.conn == nil {
		return nil
	}
	log.Debug("Close() Connection.Conn")
	err := c.conn.Close()
	c.conn = nil
	if err != nil {
		log.Warnf("FAILED Close(%s) err: %v", c.GetParams().GetDSNWithHiddenCredentials(), err)
	}
	return err
}

// connect
func (c *Connection) connect() {
	log.Debugf("Establishing connection: %s", c.GetParams().GetDSNWithHiddenCredentials())
	conn, err := databasesql.Open(c.GetParams().GetDriverName(), c.GetParams().GetDSN())
	if err != nil {
		log.Warnf("FAILED Open(%s) err: %v", c.GetParams().GetDSNWithHiddenCredentials(), err)
		return
	}

	if c.GetParams().HasMaxOpenConnections() {
		conn.SetMaxOpenConns(c.GetParams().GetMaxOpenConnections())
	}
	if c.GetParams().HasMaxIdleConnections() {
		conn.SetMaxIdleConns(c.GetParams().GetMaxIdleConnections())
	}
	if c.GetParams().HasConnectionMaxLifetime() {
		conn.SetConnMaxLifetime(c.GetParams().GetConnectionMaxLifetime())
	}
	if c.GetParams().HasConnectionMaxIdleTime() {
		conn.SetConnMaxIdleTime(c.GetParams().GetConnectionMaxIdleTime())
	}

	log.Debug("Conn Stats()")
	log.Debugf("%+v", conn.Stats())

	// Ping should be deadlined
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(c.GetParams().GetConnectTimeout()))
	defer cancel()

	if err := conn.PingContext(ctx); err != nil {
		log.Warnf("FAILED Ping(%s) err: %v", c.GetParams().GetDSNWithHiddenCredentials(), err)
		if err := conn.Close(); err != nil {
			log.Warnf("FAILED to Close conn after Ping(%s) err: %v", c.GetParams().GetDSNWithHiddenCredentials(), err)
		}
		return
	}

	c.conn = conn
}

// ensureConnected
func (c *Connection) ensureConnected() bool {
	if c == nil {
		return false
	}

	if c.conn == nil {
		log.Debugf("Need to connect to: %s", c.GetParams().GetDSNWithHiddenCredentials())
		c.connect()
	} else {
		log.Debugf("Already connected to: %s", c.GetParams().GetDSNWithHiddenCredentials())
	}

	return c.conn != nil
}

// query runs given sql query with response separated as data and error
func (c *Connection) query(sql string, args ...interface{}) (*Data, error) {
	if len(sql) == 0 {
		return nil, nil
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(c.GetParams().GetQueryTimeout()))

	if !c.ensureConnected() {
		cancel()
		s := fmt.Sprintf("FAILED connect(%s) for SQL: %s", c.GetParams().GetDSNWithHiddenCredentials(), sql)
		log.Warnf(s)
		return nil, fmt.Errorf(s)
	}

	sql = c.conn.Rebind(sql)
	rows, err := c.conn.QueryxContext(ctx, sql, args...)
	if err != nil {
		cancel()
		s := fmt.Sprintf("FAILED Query(%s) err: %v for SQL: %s", c.GetParams().GetDSNWithHiddenCredentials(), err, sql)
		log.Warnf(s)
		return nil, err
	}

	log.Debugf("conn.QueryContext():%s", sql)

	return &Data{
		ctx:        ctx,
		cancelFunc: cancel,
		Rows:       rows,
	}, nil
}

// Query runs given sql query with response as one struct which includes error (if any)
func (c *Connection) Query(sql string, args ...interface{}) *Result {
	data, err := c.query(sql, args...)
	return &Result{
		Error: err,
		Data:  data,
	}
}

// Exec runs given sql query w/o response
func (c *Connection) Exec(sql string, args ...interface{}) error {
	if len(sql) == 0 {
		return nil
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(c.GetParams().GetExecTimeout()))
	defer cancel()

	if !c.ensureConnected() {
		s := fmt.Sprintf("FAILED connect(%s) for SQL: %s", c.GetParams().GetDSNWithHiddenCredentials(), sql)
		log.Warnf(s)
		return fmt.Errorf(s)
	}

	sql = c.conn.Rebind(sql)
	_, err := c.conn.ExecContext(ctx, sql, args...)

	if err != nil {
		log.Warnf("FAILED Exec(%s) err: %v for SQL: %s", c.GetParams().GetDSNWithHiddenCredentials(), err, sql)
		return err
	}

	log.Debugf("conn.Exec():%s", sql)

	return nil
}
