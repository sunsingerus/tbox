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
	"fmt"

	databasesql "github.com/jmoiron/sqlx"

	// register db drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mailru/go-clickhouse"

	log "github.com/sirupsen/logrus"
)

// Result specifies result returned by the database - whihc encludes returned data and error code in case of an error
type Result struct {
	// Data returned by the database
	*Data
	// Error returned by the database
	Error error
}

// Close closes
func (r *Result) Close() error {
	log.Debug("Close() Result")

	if r == nil {
		return nil
	}

	err := r.Data.Close()
	r.Data = nil
	r.Error = nil
	return err
}

var (
	// ErrNilResult specifies the situation with result is nil
	ErrNilResult = fmt.Errorf("ErrNilResult")
)

// Scan scans result
func (r *Result) Scan(dest ...interface{}) error {
	if r == nil {
		return ErrNilResult
	}
	if r.Error != nil {
		return r.Error
	}
	return r.Data.Scan(dest...)
}

// ScanClose scans and closes result
func (r *Result) ScanClose(dest ...interface{}) error {
	if r == nil {
		return ErrNilResult
	}
	if r.Error != nil {
		return r.Error
	}
	err1 := r.Data.Scan(dest...)
	err2 := r.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

// Failed return true in case the result is failed
func (r *Result) Failed() bool {
	if r == nil {
		return true
	}
	if r.Error != nil {
		return true
	}
	return false
}

// Ok return true in case the result is ok
func (r *Result) Ok() bool {
	return !r.Failed()
}

// GetRows return rows returned from the database
func (r *Result) GetRows() *databasesql.Rows {
	if r == nil {
		return nil
	}
	return r.Data.GetRows()
}

// GetError returns error
func (r *Result) GetError() error {
	if r == nil {
		return nil
	}
	return r.Error
}
