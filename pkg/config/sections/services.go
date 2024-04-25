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

// ServicesConfigurator interface
type ServicesConfigurator interface {
	GetServices() items.ServicesConfigurator
}

// Interface compatibility
var _ ServicesConfigurator = Services{}

// Services specifies services list
type Services struct {
	Services *items.ServicesList `mapstructure:"services"`
}

// ServicesNormalize is a normalizer
func (c Services) ServicesNormalize() Services {
	if c.Services == nil {
		c.Services = items.NewServicesList()
	}
	return c
}

// GetServices is a getter
func (c Services) GetServices() items.ServicesConfigurator {
	return c.Services
}

// String is a stringifier
func (c Services) String() string {
	return fmt.Sprintf("Services=%s", c.Services)
}
