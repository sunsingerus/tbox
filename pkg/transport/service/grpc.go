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

	"github.com/sunsingerus/tbox/pkg/auth/service"
	"github.com/sunsingerus/tbox/pkg/config/sections"
)

type PathsTLSOAuthConfigurator interface {
	sections.PathsConfigurator
	sections.TLSConfigurator
	sections.OAuthConfigurator
	sections.OTPConfigurator
}

// GetGRPCServerOptions builds gRPC server options
func GetGRPCServerOptions(config PathsTLSOAuthConfigurator, additionalServerOptions ...grpc.ServerOption) []grpc.ServerOption {
	var opts []grpc.ServerOption

	if config.GetTLS().GetEnabled() {
		log.Infof("TLS requested")

		if transportOpts, err := setupTLS(config); err == nil {
			opts = append(opts, transportOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	}

	switch {
	case config.GetOTP().GetEnabled():
		log.Infof("OTP requested")
		if !config.GetTLS().GetEnabled() {
			log.Warn("TLS is not enabled but OTP is enabled.")
		}
		if authOpts, err := service_auth.SetupOTP(config); err == nil {
			opts = append(opts, authOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}

	case config.GetOAuth().GetEnabled():
		log.Infof("OAuth requested")
		if !config.GetTLS().GetEnabled() {
			log.Warn("TLS is not enabled but OAuth is enabled.")
		}

		if authOpts, err := service_auth.SetupOAuth(config); err == nil {
			opts = append(opts, authOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	}

	opts = append(opts, additionalServerOptions...)

	return opts
}
