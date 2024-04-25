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
	"github.com/sunsingerus/tbox/pkg/macros"
	"strings"
)

// ExitCodeIntervals specifies intervals where command's exit code lays in - negative, zero or positive.
// It can be used to make decisions on the interval of the exit code.
type ExitCodeIntervals struct {
	Negative bool `mapstructure:"negative"`
	Zero     bool `mapstructure:"zero"`
	Positive bool `mapstructure:"positive"`
}

// Check compares code with are the intervals specified and returns true in case code is covered by intervals
func (c *ExitCodeIntervals) Check(code int) bool {
	if c == nil {
		return false
	}

	switch {
	case code < 0:
		if c.Negative {
			return true
		}
	case code == 0:
		if c.Zero {
			return true
		}
	case code > 0:
		if c.Positive {
			return true
		}
	}

	return false
}

// NewExitCodeIntervalsNegative creates new negative exit code interval
func NewExitCodeIntervalsNegative() *ExitCodeIntervals {
	return &ExitCodeIntervals{
		Negative: true,
	}
}

// NewExitCodeIntervalsNegativeAndPositive creates new negative and positive exit code interval
func NewExitCodeIntervalsNegativeAndPositive() *ExitCodeIntervals {
	return &ExitCodeIntervals{
		Negative: true,
		Positive: true,
	}
}

var (
	ExitCodeIntervalsNegative            = NewExitCodeIntervalsNegative()
	ExitCodeIntervalsNegativeAndPositive = NewExitCodeIntervalsNegativeAndPositive()
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type Command struct {
	// Enabled specifies whether this command is enabled or not
	Enabled bool `mapstructure:"enabled"`
	// Workdir specifies workdir of the command
	Workdir string `mapstructure:"workdir"`
	// Env specifies ENV vars to use for a command
	Env []string `mapstructure:"env"`
	// Command specifies CLI command to launch
	Command []string `mapstructure:"command"`
	// FailExitCodeIntervals specifies in which exit code intervals command is considered to fail.
	// Allows to tune the exit code treatment and not only check for zero or non-zero value
	FailExitCodeIntervals *ExitCodeIntervals `mapstructure:"failExitCodeIntervals"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewCommand is a constructor
func NewCommand() *Command {
	return new(Command)
}

// GetEnabled is a getter
func (c *Command) GetEnabled() bool {
	if c == nil {
		return false
	}
	return c.Enabled
}

// GetWorkdir is a getter
func (c *Command) GetWorkdir() string {
	if c == nil {
		return ""
	}
	return c.Workdir
}

// GetEnv is a getter
func (c *Command) GetEnv() []string {
	if c == nil {
		return nil
	}
	return c.Env
}

// GetCommand is a getter
func (c *Command) GetCommand() []string {
	if c == nil {
		return nil
	}
	return c.Command
}

// GetCommandLine returns commdn as a line
func (c *Command) GetCommandLine() string {
	return strings.Join(c.GetCommand(), " ")
}

// ExpandCommand gets slice of strings which represents command with parameters with macros expanded
func (c *Command) ExpandCommand(macro *macros.Expander) []string {
	return macro.ExpandAll(c.GetCommand()...)
}

// ExpandCommandLine gets one string of a command with parameters with macros expanded
func (c *Command) ExpandCommandLine(macro *macros.Expander) string {
	return strings.Join(c.ExpandCommand(macro), " ")
}

// ExitCodeReportsFailure checks whether provided exit code reports program failure
func (c *Command) ExitCodeReportsFailure(code int) bool {
	if c == nil {
		return false
	}
	if c.FailExitCodeIntervals == nil {
		// No bad exit code interval explicitly specified - use default one
		return ExitCodeIntervalsNegativeAndPositive.Check(code)
	} else {
		return c.FailExitCodeIntervals.Check(code)
	}
}

// String is a stringifier
func (c *Command) String() string {
	if c == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, strings.Join(c.Command, " "))

	return b.String()
}
