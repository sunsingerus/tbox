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

	databasesql "github.com/jmoiron/sqlx"

	// register db drivers
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mailru/go-clickhouse"

	log "github.com/sirupsen/logrus"
)

// Data describes data returned from the database
type Data struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	// Rows specifies rows returned from the database
	Rows *databasesql.Rows
}

// Close closes the data returned from the database
func (d *Data) Close() error {
	log.Debug("Close() Data")

	if d == nil {
		return nil
	}

	if d.Rows != nil {
		log.Debug("Close() Data.Rows")
		err := d.Rows.Close()
		d.Rows = nil
		if err != nil {
			log.Warnf("UNABLE to close rows. err: %v", err)
			return err
		}
	}

	if d.cancelFunc != nil {
		log.Debug("Close() Data.Cancel")
		d.cancelFunc()
		d.cancelFunc = nil
	}

	if d.ctx != nil {
		log.Debug("Close() Data.Ctx")
		d.ctx = nil
	}

	return nil
}

// GetRows return rows returned from the database
func (d *Data) GetRows() *databasesql.Rows {
	if d == nil {
		return nil
	}
	return d.Rows
}

var (
	// ErrNilData specifies the situation with result data is nil
	ErrNilData = fmt.Errorf("ErrNilData")
	// ErrNilRows specifies the situation with result rows is nil
	ErrNilRows = fmt.Errorf("ErrNilRows")
	// ErrEmptyRows specifies the situation with result no rows available to scan from
	ErrEmptyRows = fmt.Errorf("ErrEmptyRows")
)

// Scan scans from the rows into set of variables
func (d *Data) Scan(dest ...interface{}) error {
	if d == nil {
		return ErrNilData
	}
	if d.Rows == nil {
		return ErrNilRows
	}

	for d.Rows.Next() {
		if err := d.Rows.Scan(dest...); err != nil {
			log.Error(err)
			return err
		}
		return nil
	}
	return ErrEmptyRows
}

func (d *Data) AllFn(extractor func(*databasesql.Rows) (interface{}, error)) (result []interface{}, err error) {
	if d == nil {
		return
	}
	if d.Rows == nil {
		return
	}
	for d.Rows.Next() {
		if item, err := extractor(d.Rows); err == nil {
			result = append(result, item)
		} else {
			log.Error(err)
		}
	}
	return result, nil
}
