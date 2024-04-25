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

package items

import (
	"bytes"
	"fmt"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type Log struct {
	Level     string `mapstructure:"level"`
	Formatter string `mapstructure:"format"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewLog is a constructor
func NewLog() *Log {
	return new(Log)
}

// GetLevel is a getter
func (l *Log) GetLevel() string {
	if l == nil {
		return ""
	}
	return l.Level
}

// GetFormatter is a getter
func (l *Log) GetFormatter() string {
	if l == nil {
		return ""
	}
	return l.Formatter
}

// String is a stringifier
func (l *Log) String() string {
	if l == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Level: %v\n", l.Level)
	_, _ = fmt.Fprintf(b, "Formatter: %v\n", l.Formatter)

	return b.String()
}
