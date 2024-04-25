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

package sections

import (
	"fmt"
	"strings"

	"github.com/sunsingerus/tbox/pkg/config/items"
	"github.com/sunsingerus/tbox/pkg/macros"
)

// CommandConfigurator
type CommandConfigurator interface {
	GetCommandWorkdir() string
	GetCommandEnv() []string
	GetCommandLines() []string
	GetCommand() string
	ParseCommandLines(*macros.Expander) []string
	ParseCommand(*macros.Expander) string
}

// Interface compatibility
var _ CommandConfigurator = Command{}

// Command
type Command struct {
	Command *items.Command `mapstructure:"command"`
}

// CommandNormalize
func (c Command) CommandNormalize() Command {
	if c.Command == nil {
		c.Command = items.NewCommand()
	}
	return c
}

// GetCommandWorkdir
func (c Command) GetCommandWorkdir() string {
	return c.Command.GetWorkdir()
}

// GetCommandEnv
func (c Command) GetCommandEnv() []string {
	return c.Command.GetEnv()
}

// GetCommandLines
func (c Command) GetCommandLines() []string {
	return c.Command.GetCommand()
}

// GetCommand
func (c Command) GetCommand() string {
	return strings.Join(c.GetCommandLines(), " ")
}

// ParseCommandLines
func (c Command) ParseCommandLines(macro *macros.Expander) []string {
	return macro.ExpandAll(c.GetCommandLines()...)
}

// ParseCommand
func (c Command) ParseCommand(macro *macros.Expander) string {
	return strings.Join(c.ParseCommandLines(macro), " ")
}

// String
func (c Command) String() string {
	return fmt.Sprintf("Command=%s", c.Command)
}
