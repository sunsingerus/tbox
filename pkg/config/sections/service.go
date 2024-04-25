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

// ServiceConfigurator
type ServiceConfigurator interface {
	GetServiceAddress() string
}

// Interface compatibility
var _ ServiceConfigurator = Service{}

// Service
type Service struct {
	Service *items.Service `mapstructure:"service"`
}

// ServiceNormalize
func (c Service) ServiceNormalize() Service {
	if c.Service == nil {
		c.Service = items.NewService()
	}
	return c
}

// GetServiceAddress
func (c Service) GetServiceAddress() string {
	return c.Service.GetAddress()
}

// String
func (c Service) String() string {
	return fmt.Sprintf("Service=%s", c.Service)
}
