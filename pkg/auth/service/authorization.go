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
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	jwks2 "github.com/sunsingerus/tbox/pkg/auth/jwks"
)

// Errors
var (
	ErrMissingMetadata      = status.Errorf(codes.InvalidArgument, "No metadata provided")
	ErrMissingAuthorization = status.Errorf(codes.Unauthenticated, "No authorization data or header provided")
	ErrMissingToken         = status.Errorf(codes.Unauthenticated, "No authorization token provided")
	ErrMissingBearer        = status.Errorf(codes.Unauthenticated, "No bearer token provided within authorization token")
	ErrInvalidToken         = status.Errorf(codes.Unauthenticated, "Invalid token")
	ErrInvalidClaims        = status.Errorf(codes.Unauthenticated, "Invalid claims")

	ErrInvalidMapClaims = fmt.Errorf("unable to map claims")
)

var (
	// jwks specifies public key set to be used by the server for JWT verification
	jwks *jwks2.JWKS
)

// dumpMetadata
func dumpMetadata(md metadata.MD) {
	t := map[string][]string(md)
	dump(t)
}

// dumpHeader
func dumpHeader(header http.Header) {
	t := map[string][]string(header)
	dump(t)
}

func dump(m map[string][]string) {
	log.Infof("Dump Header/Metadata ---")
	// Metadata is a map[string][]string
	for key, value := range m {
		log.Infof("[%s]=", key)
		for _, str := range value {
			log.Infof("    %s", str)
		}
	}
	log.Infof("End Dump Header/Metadata ---")
}
