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
type ServicesMap map[string]*Service

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT

var _ ServicesConfigurator = &ServicesMap{}

// NewServicesMap creates new services map
func NewServicesMap() *ServicesMap {
	m := new(ServicesMap)
	*m = make(map[string]*Service)
	return m
}

// GetList return list
func (s *ServicesMap) GetList() []ServiceConfigurator {
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
func (s *ServicesMap) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

// Get gets a service
func (s *ServicesMap) Get(name string) ServiceConfigurator {
	if s == nil {
		return nil
	}
	if *s == nil {
		return nil
	}

	if service, ok := (*s)[name]; ok {
		// Assign service name in case it is not provided
		if service.GetName() == "" {
			if service != nil {
				service.Name = name
			}
		}
		return service
	}

	return nil
}

// String is a stringifier
func (s *ServicesMap) String() string {
	if s == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	for _, service := range s.GetList() {
		_, _ = fmt.Fprintf(b, "%s\n", service)
	}

	return b.String()
}
