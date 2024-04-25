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

package jwks

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/sunsingerus/tbox/pkg/auth/pem"
)

// JWK specifies one JSON Web Key
type JWK struct {
	// Alg specifies cryptographic algorithm used with the key. Ex.: "RS256"
	Alg string `json:"alg" yaml:"alg"`
	// Kty specifies family of cryptographic algorithms used with the key. Ex.: "RSA"
	Kty string `json:"kty" yaml:"kty"`
	// Use specifies how the key was meant to be used; Ex.: "sig" represents the signature.
	Use string `json:"use" yaml:"use"`
	// N specifies modulus for the RSA public key. Ex.: "vY07WxvavajnrJRe6...."
	N string `json:"n" yaml:"n"`
	// E specifies exponent for the RSA public key. Ex.: "AQAB"
	E string `json:"e" yaml:"e"`
	// Kid specifies identifier of the key. Ex.: "M_XXXX-n"
	Kid string `json:"kid" yaml:"kid"`
	// X5t specifies thumbprint of the x.509 cert (SHA-1 thumbprint). Ex.: "kXXNB-yYYYYYt"
	X5t string `json:"x5t" yaml:"x5t"`
	// X5c specifies x.509 certificate chain.
	// X5c[0] is the certificate to use for token verification
	// X5c[1:] [OPTIONAL, not necessary to be included] can be used to verify X5c[0].
	X5c []string `json:"x5c" yaml:"x5c"`

	File string `json:"file,omitempty" yaml:"file,omitempty"`
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	// PublicKey is a parsed public key extracted from X5c chain
	PublicKey *rsa.PublicKey
}

var (
	ErrEmptyJWK = fmt.Errorf("specified JWK is empty")
)

// NewJWKFromString
func NewJWKFromString(cert string) (*JWK, error) {
	jwk := &JWK{
		X5c: []string{
			pem.TrimPublicKey(cert),
		},
	}

	err := jwk.Parse()
	if err != nil {
		return nil, err
	}
	return jwk, nil
}

// NewJWKFromBytes
func NewJWKFromBytes(bytes []byte) (*JWK, error) {
	return NewJWKFromString(string(bytes))
}

// NewJWKFromFile
func NewJWKFromFile(filename string) (*JWK, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read JWK from file '%s': %v", filename, err)
	}

	return NewJWKFromBytes(bytes)
}

// Parse
func (jwk *JWK) Parse() error {
	if jwk == nil {
		return ErrEmptyJWK
	}
	if len(jwk.X5c) < 1 {
		return ErrEmptyJWK
	}
	key, err := pem.ParseRSAPublicKeyFromPEM(pem.EnsurePublicKeyPEMFormat(jwk.X5c[0]))
	if err != nil {
		return err
	}
	jwk.PublicKey = key
	return nil
}

// ReadIn
func (jwk *JWK) ReadIn() *JWK {
	if jwk == nil {
		return nil
	}
	if jwk.File != "" {
		if new, err := NewJWKFromFile(jwk.File); (new != nil) && (err == nil) {
			new.FillNonEmptyFrom(jwk)
			return new
		}
	}
	if jwk.Data != "" {
		if new, err := NewJWKFromString(jwk.Data); (new != nil) && (err == nil) {
			new.FillNonEmptyFrom(jwk)
			return new
		}
	}
	return jwk
}

// FillNonEmptyFrom
func (jwk *JWK) FillNonEmptyFrom(src *JWK) *JWK {
	if (jwk.Alg == "") && (src.Alg != "") {
		jwk.Alg = src.Alg
	}
	if (jwk.Kty == "") && (src.Kty != "") {
		jwk.Kty = src.Kty
	}
	if (jwk.Use == "") && (src.Use != "") {
		jwk.Use = src.Use
	}
	if (jwk.N == "") && (src.N != "") {
		jwk.N = src.N
	}
	if (jwk.E == "") && (src.E != "") {
		jwk.E = src.E
	}
	if (jwk.Kid == "") && (src.Kid != "") {
		jwk.Kid = src.Kid
	}
	if (jwk.X5t == "") && (src.X5t != "") {
		jwk.X5t = src.X5t
	}
	if (jwk.X5c == nil) && (src.X5c != nil) {
		jwk.X5c = src.X5c
	}

	if (jwk.File == "") && (src.File != "") {
		jwk.File = src.File
	}
	if (jwk.Data == "") && (src.Data != "") {
		jwk.Data = src.Data
	}

	if (jwk.PublicKey == nil) && (src.PublicKey != nil) {
		jwk.PublicKey = src.PublicKey
	}
	return jwk
}

// String
func (jwk *JWK) String() string {
	return "jwk string rep"
}
