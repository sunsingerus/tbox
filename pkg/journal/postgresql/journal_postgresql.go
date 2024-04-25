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

package postgresql

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/journal"
	"github.com/sunsingerus/tbox/pkg/journal/adapters/postgresql"
)

// JournalPostgreSQL
type JournalPostgreSQL struct {
	journal.BaseJournal
}

// Validate interface compatibility
var _ journal.Journaller = &JournalPostgreSQL{}

// NewJournalPostgreSQLConfig
func NewJournalPostgreSQLConfig(cfg sections.PostgreSQLConfigurator, endpointID int32, endpointInstanceID *common.UUID) (*JournalPostgreSQL, error) {
	dsn := cfg.GetPostgreSQLDSN()
	return NewJournalPostgreSQL(dsn, endpointID, endpointInstanceID)
}

// NewJournalPostgreSQL
func NewJournalPostgreSQL(dsn string, endpointID int32, endpointInstanceID *common.UUID) (*JournalPostgreSQL, error) {
	if dsn == "" {
		str := "PostgreSQL address in Config is empty"
		log.Errorf(str)
		return nil, fmt.Errorf(str)
	}
	adapter, err := postgresql.NewAdapter(dsn)
	if err != nil {
		return nil, err
	}
	journal, err := journal.NewBaseJournal(endpointID, endpointInstanceID, adapter)
	if err != nil {
		return nil, err
	}
	return &JournalPostgreSQL{
		BaseJournal: *journal,
	}, nil
}
