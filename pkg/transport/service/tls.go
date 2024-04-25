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

package service_transport

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"

	"github.com/sunsingerus/tbox/pkg/config/items"
	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/devcerts"
)

type TLSPathsConfigurator interface {
	sections.PathsConfigurator
	sections.TLSConfigurator
}

func setupTLS(config TLSPathsConfigurator) ([]grpc.ServerOption, error) {
	certFile := config.GetTLS().GetPublicCertFile()
	if certFile == "" {
		certFile = devcerts.Path("service.pem")
		if _, err := os.Stat(certFile); err != nil {
			certFile = config.GetPaths().GetFile("service.pem", "tls", items.PathsOptsRebaseOnCWD)
		}
	} else {
		if _, err := os.Stat(certFile); err != nil {
			certFile = config.GetPaths().GetFile(certFile, "tls", items.PathsOptsRebaseOnCWD)
		}
	}
	keyFile := config.GetTLS().GetPrivateKeyFile()
	if keyFile == "" {
		keyFile = devcerts.Path("service.key")
		if _, err := os.Stat(keyFile); err != nil {
			keyFile = config.GetPaths().GetFile("service.key", "tls", items.PathsOptsRebaseOnCWD)
		}
	} else {
		if _, err := os.Stat(keyFile); err != nil {
			keyFile = config.GetPaths().GetFile(keyFile, "tls", items.PathsOptsRebaseOnCWD)
		}
	}

	// TransportCredentials can be created by two ways
	// 1. Directly from files via NewServerTLSFromFile()
	// 2. Or through intermediate Certificate

	// Create TransportCredentials directly from files
	transportCredentials, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	// Create TransportCredentials through intermediate Certificate
	// needs "crypto/tls"
	// cert, err := tls.LoadX509KeyPair(testdata.Path("server1.pem"), testdata.Path("server1.key"))
	// transportCredentials := credentials.NewServerTLSFromCert(&cert)

	if err != nil {
		log.Fatalf("failed to generate credentials %v", err)
	}

	log.Infof("enabling TLS with cert=%s", certFile)
	log.Infof("enabling TLS with key =%s", keyFile)

	opts := []grpc.ServerOption{
		// Enable TLS transport for connections
		grpc.Creds(transportCredentials),
	}

	return opts, nil
}
