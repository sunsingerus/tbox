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

package client_auth

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"github.com/sunsingerus/tbox/pkg/config/sections"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// SetupOTP builds []grpc.DialOption with Custom Token support from provided configuration
func SetupOTP(config sections.OTPConfigurator) ([]grpc.DialOption, error) {
	log.Infof("Using OTP params:\nServer:%s\nClient:%s\n", config.GetOTP().GetServer(), config.GetOTP().GetClient())

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(NewOTPPerRPCCredentialsImplementation(config.GetOTP().GetServer(), config.GetOTP().GetClient())),
	}
	return opts, nil
}

// OTPPerRPCCredentialsImplementation specifies custom tokens credentials
// It has to export PerRPCCredentials interface
type OTPPerRPCCredentialsImplementation struct {
	server string
	client string
	hash   string
}

// Check interface compatibility
var _ credentials.PerRPCCredentials = OTPPerRPCCredentialsImplementation{}

// NewOTPPerRPCCredentialsImplementation creates new custom token credentials implementation
func NewOTPPerRPCCredentialsImplementation(server string, client string) OTPPerRPCCredentialsImplementation {
	O := OTPPerRPCCredentialsImplementation{
		server: server,
		client: client,
	}
	s := md5.Sum([]byte(O.client + "+" + O.server))
	b := bytes.Buffer{}
	base64.NewEncoder(base64.StdEncoding, &b).Write(s[:])
	O.hash = b.String()
	return O
}

// GetRequestMetadata is function of credentials.PerRPCCredentials
func (O OTPPerRPCCredentialsImplementation) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	h := map[string]string{
		"Authorization": "Basic " + O.hash,
	}
	return h, nil
}

// RequireTransportSecurity is function of credentials.PerRPCCredentials
func (O OTPPerRPCCredentialsImplementation) RequireTransportSecurity() bool {
	return true
}
