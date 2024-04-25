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

// ManticoreConfigurator
type ManticoreConfigurator interface {
	GetManticoreHostname() string
	GetManticorePort() int
	GetManticoreDSN() string
}

// Interface compatibility
var _ ManticoreConfigurator = Manticore{}

// Manticore
type Manticore struct {
	Manticore *items.Manticore `mapstructure:"manticore"`
}

// ManticoreNormalize
func (c Manticore) ManticoreNormalize() Manticore {
	if c.Manticore == nil {
		c.Manticore = items.NewManticore()
	}
	return c
}

// GetManticoreHostname
func (c Manticore) GetManticoreHostname() string {
	return c.Manticore.GetHostname()
}

// GetManticorePort
func (c Manticore) GetManticorePort() int {
	return c.Manticore.GetPort()
}

// GetManticoreDSN
func (c Manticore) GetManticoreDSN() string {
	return c.Manticore.GetDSN()
}

// String
func (c Manticore) String() string {
	return fmt.Sprintf("Manticore=%s", c.Manticore)
}
