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

package pem

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

const (
	beginPublicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	endPublicKeySuffix   = "-----END PUBLIC KEY-----"
)

var (
	ErrEnclosed = fmt.Errorf("public key must be enclosed in %s %s", beginPublicKeyPrefix, endPublicKeySuffix)
)

// PEM (Privacy-Enhanced Mail) is a text format for storing cryptographic keys, certificates, etc...
// Described in rfc7468
// https://datatracker.ietf.org/doc/html/rfc7468
// Main idea is the following - enclose body in
// -----BEGIN `label` -----
// -----END `label` -----
// where `label` can be one of:
// 1. CERTIFICATE
// 2. PRIVATE KEY
// 3. PUBLIC KEY
// 4. CERTIFICATE REQUEST
// etc...

// EnsurePublicKeyPEMFormat ensures PK is wrapped into -----BEGIN PK ----- and -----END PK-----
func EnsurePublicKeyPEMFormat(certStringOrBytes interface{}) []byte {
	cert := ""
	switch t := certStringOrBytes.(type) {
	case string:
		cert = t
	case []byte:
		cert = string(t)
	}
	cert = trim(cert)
	if IsPublicKeyPEMFormat(cert) {
		// Already enclosed, PEM format is in place
		return []byte(cert)
	}
	return []byte(fmt.Sprintf("%s\n%s\n%s", beginPublicKeyPrefix, cert, endPublicKeySuffix))
}

// ParseRSAPublicKeyFromPEM parses PEM public key into struct
func ParseRSAPublicKeyFromPEM(pem []byte) (*rsa.PublicKey, error) {
	// Ensure PEM is reasonable

	if err := CheckPublicKeyPEMFormat(pem); err != nil {
		return nil, err
	}

	// Parse RSA Public Key
	return jwt.ParseRSAPublicKeyFromPEM(pem)
}

// IsPublicKeyPEMFormat checks whether data has correct PEM format
// This function is a convenience wrapper for CheckPublicKeyPEMFormat
func IsPublicKeyPEMFormat(certStringOrBytes interface{}) bool {
	return CheckPublicKeyPEMFormat(certStringOrBytes) == nil
}

// CheckPublicKeyPEMFormat checks whether PEM format has any errors (properly formatted)
func CheckPublicKeyPEMFormat(certStringOrBytes interface{}) error {
	cert := ""
	switch t := certStringOrBytes.(type) {
	case string:
		cert = t
	case []byte:
		cert = string(t)
	}
	return checkPublicKeyPEMFormat(cert)
}

// checkPublicKeyPEMFormat is a typed string check
func checkPublicKeyPEMFormat(cert string) error {
	trimmed := trim(cert)
	// Key must be enclosed into start/stop tags
	if !strings.HasPrefix(trimmed, beginPublicKeyPrefix) || !strings.HasSuffix(trimmed, endPublicKeySuffix) {
		return ErrEnclosed
	}
	return nil
}

// TrimPublicKey trims PEM prefix/suffix from Public Key
func TrimPublicKey(cert string) string {
	cert = trim(cert)
	cert = strings.TrimPrefix(cert, beginPublicKeyPrefix)
	cert = trim(cert)
	cert = strings.TrimSuffix(cert, endPublicKeySuffix)
	return trim(cert)
}

// trim newlines from the end of the key
func trim(cert string) string {
	return strings.TrimRight(cert, "\r\n")
}
