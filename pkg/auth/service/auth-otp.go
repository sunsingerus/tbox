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
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/sunsingerus/tbox/pkg/config/sections"
)

// SetupOTP prepares gRPC server options for OTP auth
func SetupOTP(config sections.OAuthConfigurator) ([]grpc.ServerOption, error) {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(otpUnaryInterceptor),
		grpc.ChainStreamInterceptor(otpStreamInterceptor),
	}

	return opts, nil
}

// otpUnaryInterceptor is a default interceptor for OTP
func otpUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}
	dumpMetadata(md)
	log.Infof("OTP AUTH FAKE OK : token=%s", md["Authorization"])
	return handler(ctx, req)
}

// otpStreamInterceptor is a default interceptor for OTP
func otpStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := ss.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ErrMissingMetadata
	}
	dumpMetadata(md)
	log.Infof("OTP AUTH FAKE OK : token=%s", md["Authorization"])
	return handler(srv, ss)
}
