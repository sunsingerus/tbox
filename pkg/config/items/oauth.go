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
	"net/url"
)

// OAuthConfigurator
type OAuthConfigurator interface {
	GetEnabled() bool
	GetClientID() string
	GetClientSecret() string
	GetTokenURL() string
	GetEndpointParams() url.Values
	GetRegisterURL() string
	GetInitialAccessToken() string
	GetJWT() *JWT
}

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type OAuth struct {
	Enabled bool `mapstructure:"enabled"`

	// OAuth Login section

	// ClientID is the application's identifier.
	ClientID string `mapstructure:"client-id"`
	// ClientSecret is the application's secret.
	ClientSecret string `mapstructure:"client-secret"`
	// TokenURL is the identity server's token endpoint URL, where to send token request.
	TokenURL string `mapstructure:"token-url"`
	// EndpointParams are additional parameters for requests to the token endpoint.
	// Such as:
	// Unique identifier for the API. This value will be used as the audience parameter on authorization calls.
	//    audience:
	//      - a.b.c
	EndpointParams url.Values `mapstructure:"endpoint-params"`

	// OAuth Register section

	RegisterURL        string `mapstructure:"register-url"`
	InitialAccessToken string `mapstructure:"initial-access-token"`

	// JWT section
	JWT *JWT `mapstructure:"jwt"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

var _ OAuthConfigurator = &OAuth{}

// NewOAuth is a condtructor
func NewOAuth() *OAuth {
	return new(OAuth)
}

// GetEnabled is a getter
func (o *OAuth) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

// GetClientID is a getter
func (o *OAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

// GetClientSecret is a getter
func (o *OAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

// GetTokenURL is a getter
func (o *OAuth) GetTokenURL() string {
	if o == nil {
		return ""
	}
	return o.TokenURL
}

// GetEndpointParams is a getter
func (o *OAuth) GetEndpointParams() url.Values {
	if o == nil {
		return nil
	}
	return o.EndpointParams
}

// GetRegisterURL is a getter
func (o *OAuth) GetRegisterURL() string {
	if o == nil {
		return ""
	}
	return o.RegisterURL
}

// GetInitialAccessToken is a getter
func (o *OAuth) GetInitialAccessToken() string {
	if o == nil {
		return ""
	}
	return o.InitialAccessToken
}

// GetJWT is a getter
func (o *OAuth) GetJWT() *JWT {
	if o == nil {
		return nil
	}
	return o.JWT
}

// String is a stringifier
func (o *OAuth) String() string {
	if o == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", o.Enabled)
	_, _ = fmt.Fprintf(b, "ClientID: %v\n", o.ClientID)
	_, _ = fmt.Fprintf(b, "ClientSecret: %v\n", o.ClientSecret)
	_, _ = fmt.Fprintf(b, "TokenURL: %v\n", o.TokenURL)
	_, _ = fmt.Fprintf(b, "RegisterURL: %v\n", o.RegisterURL)
	_, _ = fmt.Fprintf(b, "InitialAccessToken: %v\n", o.InitialAccessToken)
	_, _ = fmt.Fprintf(b, "JWT: %v\n", o.JWT)

	return b.String()
}
