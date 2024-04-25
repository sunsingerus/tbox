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
	"strconv"
	"strings"
)

// Used to build DSN w/o username:password pair. Used for writing logs, etc...
const (
	usernameReplacer = "*"
	passwordReplacer = "*"
)

// ClickHouse DSN should be like http://user:password@host:8123/database
const (
	dsnURLPatternClickHouse           = "http://%s%s:%s/%s"
	dsnURLPatternSimplifiedClickHouse = "http://%s%s"
)

// MySQL DSN should be like user:password@tcp(host:3306)/database
const (
	dsnURLPatternMySQL           = "%stcp(%s:%s)/%s"
	dsnURLPatternSimplifiedMySQL = "%s%s"
)

// PostgreSQL DSN should be like postgres://user:password@host:5432/database
const (
	dsnURLPatternPostgreSQL           = "postgres://%s%s:%s/%s"
	dsnURLPatternSimplifiedPostgreSQL = "%s%s"
)

// Used to build username:password@ component of the DSN
const (
	dsnUsernamePasswordPairPattern             = "%s:%s@"
	dsnUsernamePasswordPairUsernameOnlyPattern = "%s@"
)

// makeUserPassPair makes "username:password@" part for DSN
func makeUserPassPair(username, password string, hidden bool) string {

	// In case of hidden username+password pair we'd just return replacement
	if hidden {
		return fmt.Sprintf(dsnUsernamePasswordPairPattern, usernameReplacer, passwordReplacer)
	}

	// We may have neither username nor password
	if username == "" && password == "" {
		return ""
	}

	// Password may be omitted
	if password == "" {
		return fmt.Sprintf(dsnUsernamePasswordPairUsernameOnlyPattern, username)
	}

	// Expecting both username and password to be in place
	return fmt.Sprintf(dsnUsernamePasswordPairPattern, username, password)
}

func makeDSN(driverName, username, password string, hostname string, port int, database string, hideCredentials bool) string {
	switch driverName {
	case "mysql":
		return makeDSNMySQL(username, password, hostname, port, database, hideCredentials)
	case "pgx":
		return makeDSNPostgreSQL(username, password, hostname, port, database, hideCredentials)
	case "clickhouse":
		return makeDSNClickHouse(username, password, hostname, port, database, hideCredentials)
	}
	return ""
}

// makeDSNClickHouse makes ClickHouse DSN
func makeDSNClickHouse(username, password string, hostname string, port int, database string, hideCredentials bool) string {
	return fmt.Sprintf(
		dsnURLPatternClickHouse,
		makeUserPassPair(username, password, hideCredentials),
		hostname,
		strconv.Itoa(port),
		database,
	)
}

// makeDSNMySQL makes MySQL DSN
func makeDSNMySQL(username, password string, hostname string, port int, database string, hideCredentials bool) string {
	return fmt.Sprintf(
		dsnURLPatternMySQL,
		makeUserPassPair(username, password, hideCredentials),
		hostname,
		strconv.Itoa(port),
		database,
	)
}

// makeDSNPostgreSQL makes PostgreSQL DSN
func makeDSNPostgreSQL(username, password string, hostname string, port int, database string, hideCredentials bool) string {
	return fmt.Sprintf(
		dsnURLPatternPostgreSQL,
		makeUserPassPair(username, password, hideCredentials),
		hostname,
		strconv.Itoa(port),
		database,
	)
}

func hideCredentialsDSN(driverName, dsn string) string {
	switch driverName {
	case "mysql":
		return hideCredentialsDSNMySQL(dsn)
	case "pgx":
		return hideCredentialsDSNPostgreSQL(dsn)
	case "clickhouse":
		return hideCredentialsDSNClickHouse(dsn)
	}
	return ""
}

// hideCredentialsDSNClickHouse hides credentials from the existing DSN.
// This function should be used in case it is not possible to build DSN from components (username, password, etc...),
// but we have ready DSN only, and thus, need to rebuild it.
func hideCredentialsDSNClickHouse(dsn string) string {
	// Find last @ in http://user:password@host:8123/database
	i := strings.LastIndex(dsn, "@")
	if i < 0 {
		// Not found username/password pair
		// Return as it is
		return dsn
	}

	// Found username/password pair, extract main part (after @),
	// which would be "host:8123/database"
	main := dsn[i+1:]
	// Rebuild DSN
	return fmt.Sprintf(
		dsnURLPatternSimplifiedClickHouse,
		makeUserPassPair("", "", true),
		main,
	)
}

// hideCredentialsDSNMySQL hides credentials from the existing DSN.
// This function should be used in case it is not possible to build DSN from components (username, password, etc...),
// but we have ready DSN only, and thus, need to rebuild it.
func hideCredentialsDSNMySQL(dsn string) string {
	// Find last @ in user:password@tcp(host:3306)/database
	i := strings.LastIndex(dsn, "@")
	if i < 0 {
		// Not found username/password pair
		// Return as it is
		return dsn
	}

	// Found username/password pair, extract main part (after @),
	// which would be "tcp(host:3306)/database"
	main := dsn[i+1:]
	// Rebuild DSN
	return fmt.Sprintf(
		dsnURLPatternSimplifiedMySQL,
		makeUserPassPair("", "", true),
		main,
	)
}

// hideCredentialsDSNPostgreSQL hides credentials from the existing DSN.
// This function should be used in case it is not possible to build DSN from components (username, password, etc...),
// but we have ready DSN only, and thus, need to rebuild it.
func hideCredentialsDSNPostgreSQL(dsn string) string {
	// Find last @ in postgres://user:password@host:5432/database
	i := strings.LastIndex(dsn, "@")
	if i < 0 {
		// Not found username/password pair
		// Return as it is
		return dsn
	}

	// Found username/password pair, extract main part (after @),
	// which would be "host:5432/database"
	main := dsn[i+1:]
	// Rebuild DSN
	return fmt.Sprintf(
		dsnURLPatternSimplifiedPostgreSQL,
		makeUserPassPair("", "", true),
		main,
	)
}
