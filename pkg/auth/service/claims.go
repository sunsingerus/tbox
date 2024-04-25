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

package service_auth

import (
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

// ScopeClaims define scope token.
// It is an Access Token, which provides scope (set of permissions) which this token (claim) provides.
// Scope is a synonym to permission(s), both terms are usable.
// Token example 1:
//
//	{
//	  aud: "https://aud",
//	  exp: 1630351936,
//	  iat: 1630265536,
//	  iss: "https://<tenant>.auth0.com/",
//	  azp: "auth0 ClientID of the application",
//	  gty: "client-credentials"
//	  sub: "sub@clients",
//	  scope: "permission-scope-read",
//	}
//
// Token example 2:
//
//	{
//	  iss: "https://<tenant>.auth0.com/",
//	  sub: "auth0|user id goes here 3e65",
//	  aud: [
//	    "audience",
//	    "https://auth0.com/userinfo"
//	  ],
//	  iat: 1637762859,
//	  exp: 1637849259,
//	  azp: "auth0 ClientID of the application",
//	  scope: "openid profile email"
//	}
//
// Where:
//
//		 aud: Audience
//		 exp: Expires At
//		 jti: Id
//		 iat: Issued At
//		 iss: Issuer
//		 nbf: Not Before
//		 sub: Subject
//	  azp: Authorized Parties (OIDC claims)
type ScopeClaims struct {
	jwt.StandardClaims
	// Scope is a synonym to permission(s). Set of space-separated items.
	Scope string `json:"scope"`
}

// Valid checks whether claims are valid.
func (c ScopeClaims) Valid() error {
	// TODO scope field is not validated ATM
	return c.StandardClaims.Valid()
}

// Dump logs claims
func (c ScopeClaims) Dump() {
	log.Infof("Customized Claims:")
	log.Infof("Audience  %s", c.Audience)
	log.Infof("ExpiresAt %d", c.ExpiresAt)
	log.Infof("Id        %s", c.Id)
	log.Infof("IssuedAt  %d", c.IssuedAt)
	log.Infof("Issuer    %s", c.Issuer)
	log.Infof("NotBefore %d", c.NotBefore)
	log.Infof("Subject   %s", c.Subject)
	log.Infof("Scope     %s", c.Scope)
}
