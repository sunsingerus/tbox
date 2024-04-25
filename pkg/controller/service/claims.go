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

package controller_service

import (
	"context"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/auth/service"
)

// ClaimsExtractorFunction is a function, used to extract claims from incoming gRPC request context.
// Claims enclosed in the context can vary, depending on what exactly client send.
// Server has to understand what claims are sent, that's why we need to have claims extractor configurable.
type ClaimsExtractorFunction = func(context.Context) jwt.Claims

// ClaimsExtractor provides function to extract claims.
// This function can be provided by user, however, there are the following predefined functions ATM:
// 1. ClaimsExtractorMapClaims extracts claims as jwt.MapClaims from context
// 2. ClaimsExtractorScopeClaims extracts claims as service_auth.ScopeClaims from context
var ClaimsExtractor ClaimsExtractorFunction = ClaimsExtractorMapClaims

// ExtractClaims in an interface function expected to be used by user after ClaimsExtractor was set up.
// It mainly wraps if into one-liner.
func ExtractClaims(ctx context.Context) jwt.Claims {
	if ClaimsExtractor == nil {
		return nil
	}
	return ClaimsExtractor(ctx)
}

// Type verification
var (
	_ ClaimsExtractorFunction = ClaimsExtractorMapClaims
)

// ClaimsExtractorMapClaims extracts claims as jwt.MapClaims from context
func ClaimsExtractorMapClaims(ctx context.Context) jwt.Claims {
	claims, err := service_auth.GetMapClaimsGRPC(ctx)
	if err != nil {
		log.Warnf("unable to get claims with err: %v", err)
		return nil
	}

	log.Infof("ClaimsExtractorMapClaims() Claims:")
	for name, value := range claims {
		log.Infof("%s: %v", name, value)
	}

	return claims
}
