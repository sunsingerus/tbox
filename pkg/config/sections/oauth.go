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

// OAuthConfigurator
type OAuthConfigurator interface {
	GetOAuth() items.OAuthConfigurator
}

// Interface compatibility
var _ OAuthConfigurator = OAuth{}

// OAuth
type OAuth struct {
	OAuth *items.OAuth `mapstructure:"oauth"`
}

// OAuthNormalize is a normalizer
func (c OAuth) OAuthNormalize() OAuth {
	if c.OAuth == nil {
		c.OAuth = items.NewOAuth()
	}
	return c
}

// GetOAuth is a getter
func (c OAuth) GetOAuth() items.OAuthConfigurator {
	return c.OAuth
}

// String is a stringifier
func (c OAuth) String() string {
	return fmt.Sprintf("OAuth=%s", c.OAuth)
}
