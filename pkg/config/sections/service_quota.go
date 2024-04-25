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

// RESTServiceConfigurator
type QuotaServiceConfigurator interface {
	GetQuotaServiceAddress() string
}

// Interface compatibility
var _ QuotaServiceConfigurator = QuotaService{}

// QuotaService
type QuotaService struct {
	Service *items.Service `mapstructure:"quota"`
}

// QuotaServiceNormalize
func (c QuotaService) QuotaServiceNormalize() QuotaService {
	if c.Service == nil {
		c.Service = items.NewService()
	}
	return c
}

// GetQuotaServiceAddress
func (c QuotaService) GetQuotaServiceAddress() string {
	return c.Service.GetAddress()
}

// String
func (c QuotaService) String() string {
	return fmt.Sprintf("QuotaService=%s", c.Service)
}
