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
	"github.com/sunsingerus/tbox/pkg/config/items"
)

// NamedCommandsConfigurator
type NamedCommandsConfigurator interface {
	GetCommands() *items.NamedCommands
}

// Interface compatibility
var _ NamedCommandsConfigurator = NamedCommands{}

// NamedCommands is a named commands
type NamedCommands struct {
	Commands *items.NamedCommands `mapstructure:"commands"`
}

// NamedCommandsNormalize performs normalization process
func (c NamedCommands) NamedCommandsNormalize() NamedCommands {
	if c.Commands == nil {
		c.Commands = items.NewNamedCommands()
	}
	return c
}

// GetCommands is a getter
func (c NamedCommands) GetCommands() *items.NamedCommands {
	return c.Commands
}

// String is a stringifier
func (c NamedCommands) String() string {
	return fmt.Sprintf("Commands=%s", c.Commands)
}
