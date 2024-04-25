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

// PathsConfigurator
type PathsConfigurator interface {
	GetPaths() items.PathsConfigurator
}

// MultiPaths specifies paths list
type MultiPaths struct {
	Paths *items.MultiPaths `mapstructure:"paths"`
}

// PathsNormalize is a normalizer
func (c MultiPaths) PathsNormalize() MultiPaths {
	if c.Paths == nil {
		c.Paths = items.NewMultiPaths()
	}
	return c
}

// GetPaths is a getter
func (c MultiPaths) GetPaths() items.PathsConfigurator {
	return c.Paths
}

// String is a stringifier
func (c MultiPaths) String() string {
	return fmt.Sprintf("Paths=%s", c.Paths)
}

// Paths specifies paths list
type Paths struct {
	Paths *items.Paths `mapstructure:"paths"`
}

// PathsNormalize is a normalizer
func (c Paths) PathsNormalize() Paths {
	if c.Paths == nil {
		c.Paths = items.NewPaths()
	}
	return c
}

// GetPaths is a getter
func (c Paths) GetPaths() items.PathsConfigurator {
	return c.Paths
}

// String is a stringifier
func (c Paths) String() string {
	return fmt.Sprintf("Paths=%s", c.Paths)
}
