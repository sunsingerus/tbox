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
	"context"
	"fmt"

	"github.com/coreos/go-oidc"
	log "github.com/sirupsen/logrus"
	cc "golang.org/x/oauth2/clientcredentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/sunsingerus/tbox/pkg/config/sections"
)

// SetupOAuth builds []grpc.DialOption with OAuth2 support from provided configuration
func SetupOAuth(config sections.OAuthConfigurator) ([]grpc.DialOption, error) {
	log.Infof("Using OAuth params:\nClientID:%s\nTokenURL:%s\n", config.GetOAuth().GetClientID(), config.GetOAuth().GetTokenURL())

	// 2-legged OAuth2 flow
	oAuthFlow := &cc.Config{
		ClientID:       config.GetOAuth().GetClientID(),
		ClientSecret:   config.GetOAuth().GetClientSecret(),
		TokenURL:       config.GetOAuth().GetTokenURL(),
		EndpointParams: config.GetOAuth().GetEndpointParams(),
	}
	// Make OAuth2-native TokenSource object which can provide OAuth2 tokens
	ts := oAuthFlow.TokenSource(context.Background())

	// Let's fetch token right now, in order to check, whether token can be obtained with provided config
	if t, err := ts.Token(); err != nil {
		return nil, fmt.Errorf("unable to obtain oauth2 token: %v", err)
	} else if !t.Valid() {
		return nil, fmt.Errorf("oauth2 token is not valid at the beginning")
	}

	// Token can be fetched, may use OAuth2 TokenSource with gRPC

	// Since gRPC is not limited with OAuth2, it has its own TokenSource, which is wider than OAuth2-s
	// Yes, these name duplication is a little confusing.
	// Build gRPC TokenSource from OAuth2 TokenSource
	grpcTokenSource := oauth.TokenSource{
		TokenSource: ts,
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(grpcTokenSource),
	}

	return opts, nil
}

func qwe() {
	_, err := oidc.NewProvider(context.Background(), "providerURI")
	if err != nil {
		log.Fatal(err)
	}
}
