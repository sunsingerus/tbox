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

// NamedCommands is a set named commands. Each command has string name by which it can be accessed.
// IMPORTANT.
// IMPORTANT. Do not forget to update String() function
// IMPORTANT.
type NamedCommands map[string]*Command

// IMPORTANT.
// IMPORTANT. Do not forget to update String() function
// IMPORTANT.

// NewNamedCommands creates new named commands
func NewNamedCommands() *NamedCommands {
	commands := new(NamedCommands)
	*commands = make(map[string]*Command)
	return commands
}

// GetNames gets list of names in commands
func (c *NamedCommands) GetNames() []string {
	if c == nil {
		return nil
	}
	var res []string
	for name := range *c {
		res = append(res, name)
	}
	return res
}

// HasCommand checks whether there is a command with specified name
func (c *NamedCommands) HasCommand(name string) bool {
	return c.GetCommand(name) != nil
}

// GetCommand returns command by its name
func (c *NamedCommands) GetCommand(name string) *Command {
	if c == nil {
		return nil
	}
	if *c == nil {
		return nil
	}

	if cmd, ok := (*c)[name]; ok {
		return cmd
	}

	return nil
}

// String is a stringifier
func (c *NamedCommands) String() string {
	if c == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	for name, command := range *c {
		_, _ = fmt.Fprintf(b, "%s:%s\n", name, command)
	}

	return b.String()
}
