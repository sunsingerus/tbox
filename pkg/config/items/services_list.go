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
type ServicesList []*Service

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT

var _ ServicesConfigurator = &ServicesList{}

// NewServicesList creates new services list
func NewServicesList() *ServicesList {
	list := new(ServicesList)
	*list = make([]*Service, 0)
	return list
}

// GetList return list
func (s *ServicesList) GetList() []ServiceConfigurator {
	if s == nil {
		return nil
	}

	var list []ServiceConfigurator
	for _, service := range *s {
		list = append(list, service)
	}

	return list
}

// Len returns number of entries in the services list
func (s *ServicesList) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

// Get gets service
func (s *ServicesList) Get(name string) ServiceConfigurator {
	if s == nil {
		return nil
	}
	for _, service := range s.GetList() {
		if service.GetName() == name {
			return service
		}
	}
	return nil
}

// String is a stringifier
func (s *ServicesList) String() string {
	if s == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	for _, service := range s.GetList() {
		_, _ = fmt.Fprintf(b, "%s\n", service)
	}

	return b.String()
}
