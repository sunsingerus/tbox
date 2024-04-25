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

package clickhouse

import (
	"fmt"

	databasesql "github.com/jmoiron/sqlx"

	"github.com/MakeNowJust/heredoc"
	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/journal"
)

// Adapter adapts journal for ClickHouse storage
type Adapter struct {
	connect *databasesql.DB
}

// Validate interface compatibility
var _ journal.Adapter = &Adapter{}

// NewAdapterFromConfig creates new Adapter from config
func NewAdapterFromConfig(cfg sections.ClickHouseConfigurator) (*Adapter, error) {
	dsn := cfg.GetClickHouseDSN()
	return NewAdapter(dsn)
}

// NewAdapter creates new Adapter for DSN string
func NewAdapter(dsn string) (*Adapter, error) {
	if dsn == "" {
		str := "ClickHouse address in Config is empty"
		log.Errorf(str)
		return nil, fmt.Errorf(str)
	}

	log.Infof("connect to ClickHouse %s", dsn)

	connect, err := databasesql.Open("clickhouse", dsn)
	if err != nil {
		log.Errorf("unable to open ClickHouse err: %v", err)
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		log.Errorf("unable to ping ClickHouse. err: %v", err)
		return nil, err
	}

	return &Adapter{
		connect: connect,
	}, nil
}

// Insert inserts journal entry into ClickHouse
func (j *Adapter) Insert(entry *journal.Entry) error {
	e := NewAdapterEntry().Import(entry)
	sql := heredoc.Docf(`
		INSERT INTO api_journal (
			%s
		) VALUES (
			%s
		)
		`,
		e.Fields(),
		e.StmtParamsPlaceholder(),
	)

	sql = j.connect.Rebind(sql)

	tx, err := j.connect.Begin()
	if err != nil {
		log.Errorf("unable to begin tx. err: %v", err)
		return err
	}

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorf("unable to prepare stmt. err: %v", err)
		return err
	}

	if _, err := stmt.Exec(e.AsUntypedSlice()...); err != nil {
		log.Errorf("exec failed. err: %v", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("commit failed. err %v", err)
		return err
	}

	return nil
}

// FindAll finds entries in ClickHouse
func (j *Adapter) FindAll(entry *journal.Entry) ([]*journal.Entry, error) {
	e := NewAdapterEntryClickHouseSearch().Import(entry)
	placeholder, args := e.StmtSearchParamsPlaceholderAndArgs()
	sql := heredoc.Doc(
		fmt.Sprintf(`
			SELECT * FROM api_journal WHERE (1 == 1) %s 
			`,
			placeholder,
		),
	)

	stmt, err := j.connect.Preparex(sql)
	if err != nil {
		log.Errorf("unable to prepare stmt. err: %v", err)
		return nil, err
	}

	rows, err := stmt.Queryx(args...)
	if err != nil {
		log.Errorf("unable to query stmt. err: %v", err)
		return nil, err
	}
	defer rows.Close()

	var res []*journal.Entry
	for rows.Next() {
		ce := NewAdapterEntry()
		if err := ce.Scan(rows); err == nil {
			res = append(res, ce.Export())
		} else {
			log.Errorf("unable to scan stmt. err: %v", err)
		}
	}

	return res, nil
}
