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
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/sunsingerus/tbox/pkg/config/sections"
)

// SetupOAuth prepares gRPC server options with OAuth from config
func SetupOAuth(config sections.OAuthConfigurator) ([]grpc.ServerOption, error) {
	// TODO refactor global var usage
	// Prepare RSA public key to be used for JWT parsing
	jwks = config.GetOAuth().GetJWT().GetJWKS().ReadIn()

	// As we have public key set to parse JWT,
	// we can set up interceptors to perform server-side authorization
	opts := []grpc.ServerOption{
		// Add an interceptor for all unary RPCs.
		grpc.ChainUnaryInterceptor(oAuthUnaryInterceptor),

		// Add an interceptor for all stream RPCs.
		grpc.ChainStreamInterceptor(oAuthStreamInterceptor),
	}

	return opts, nil
}

// oAuthUnaryInterceptor is a default oAuth interceptor
// In case of failed authorization, the interceptor blocks execution of the handler and returns an error.
// type grpc.StreamClientInterceptor
func oAuthUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Infof("unaryInterceptor %s", info.FullMethod)

	// Skip authorize when GetJWT is requested
	//if info.FullMethod != "/proto.EventStoreService/GetJWT" {
	//	if err := authorize(ctx); err != nil {
	//		return nil, err
	//	}
	//}

	if err := AuthorizeGRPC(ctx); err != nil {
		log.Infof("AUTH FAILED unaryInterceptor %s %v", info.FullMethod, err.Error())
		return nil, err
	}

	log.Infof("AUTH OK unaryInterceptor %s", info.FullMethod)

	// Continue execution of handler
	return handler(ctx, req)
}

// oAuthStreamInterceptor is a default oAuth interceptor
// In case of failed authorization, the interceptor blocks execution of the handler and returns an error.
// type grpc.StreamServerInterceptor
func oAuthStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Infof("streamInterceptor %s %t %t", info.FullMethod, info.IsClientStream, info.IsServerStream)

	// Skip authorize when GetJWT is requested
	//if info.FullMethod != "/proto.EventStoreService/GetJWT" {
	//	if err := authorize(ctx); err != nil {
	//		return nil, err
	//	}
	//}

	ctx := ss.Context()
	if err := AuthorizeGRPC(ctx); err != nil {
		log.Infof("AUTH FAILED streamInterceptor %s %v", info.FullMethod, err.Error())
		return err
	}

	log.Infof("AUTH OK streamInterceptor %s %t %t", info.FullMethod, info.IsClientStream, info.IsServerStream)

	// Continue execution of handler
	return handler(srv, ss)
}

// AuthorizeGRPC performs gRPC auth
func AuthorizeGRPC(ctx context.Context) error {
	_, err := GetMapClaimsGRPC(ctx)
	return err
}

// AuthorizeHTTP performs HTTP auth
func AuthorizeHTTP(request *http.Request) error {
	_, err := GetMapClaimsHTTP(request)
	return err
}

// GetMapClaimsGRPC ensures a valid token exists within a request's metadata and authorizes the token received from Metadata
func GetMapClaimsGRPC(ctx context.Context) (jwt.MapClaims, error) {
	// Fetch metadata from request's context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}
	dumpMetadata(md)

	// Fetch authorization metadata from request's metadata
	authorization, ok := md["authorization"]
	if !ok {
		return nil, ErrMissingAuthorization
	}

	return GetBearerMapClaims(authorization)
}

// GetMapClaimsHTTP ensures a valid token exists within a request's metadata and authorizes the token received from Metadata
func GetMapClaimsHTTP(request *http.Request) (jwt.MapClaims, error) {
	if len(request.Header) == 0 {
		return nil, ErrMissingMetadata
	}
	dumpHeader(request.Header)

	// Fetch authorization metadata from request's metadata
	authorization, ok := request.Header["Authorization"]
	if !ok {
		authorization, ok = request.Header["authorization"]
		if !ok {
			return nil, ErrMissingAuthorization
		}
	}

	return GetBearerMapClaims(authorization)
}

// GetBearerMapClaims get map claims from 'Bearer XXX' token
func GetBearerMapClaims(authorizationHeader []string) (jwt.MapClaims, error) {
	claims, err := GetBearerClaims(authorizationHeader, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	typedClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidMapClaims
	}

	return typedClaims, nil
}

// GetBearerClaims fetches authorization claims from a request's authorization header
// NB: `claims` is used as an output value
// Provided `claims` is filled with values from the token.
// Thus, `claims` jwt.Claims should be writable-by-value type, such as
// a pointer to a struct, such as
// Ex.: &ScopeClaims{}
// or it can be a map,
// Ex.: jwt.MapClaims{}
// because, `claims` is filled with data, fetched from context
func GetBearerClaims(authorizationHeader []string, claims jwt.Claims) (jwt.Claims, error) {
	token, err := fetchAndVerifyJWTBearerToken(authorizationHeader, claims)
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

// fetchAndVerifyJWTBearerToken
func fetchAndVerifyJWTBearerToken(authorizationHeader []string, claims jwt.Claims) (*jwt.Token, error) {
	// Fetch bearer token (as a string) from authorization header
	token, err := fetchBearerToken(authorizationHeader)
	if err != nil {
		return nil, err
	}

	// We have token in place, let's parse the token into claims
	return parseAndVerifyToken(token, claims)
}

// fetchBearerToken fetches token (as a string) from authorization header/metadata "Bearer XXX" header.
// `header` can be provided either from HTTP headers or as gRPC authorization metadata.
// In any case it is expected to be of type []string having "Bearer XXX" as a first (0-indexed) entry.
func fetchBearerToken(authorizationHeader []string) (string, error) {
	if len(authorizationHeader) < 1 {
		return "", ErrMissingToken
	}

	// Prefix of the "Bearer XXX" line
	prefix := "Bearer "

	// Fetch token line "Bearer XXX"
	bearer := authorizationHeader[0]
	if !strings.HasPrefix(bearer, prefix) {
		return "", ErrMissingBearer
	}

	// Fetch bearer token itself - trim prefix from "Bearer XXX"
	token := strings.TrimPrefix(bearer, prefix)
	if len(token) < 1 {
		return "", ErrMissingBearer
	}
	log.Infof("Bearer %s", token)

	return token, nil
}

// getTokenVerificationKey gets key/cert used for JWT verification
func getTokenVerificationKey(token *jwt.Token) (interface{}, error) {
	// Variable 'token *jwt.Token' contains parsed, but unverified Token.
	// Function must return the key for verification of the specified token.
	// Receiving parsed token, allows us to use properties in the Header (such as `kid`) to identify which key to use.
	// For examples, to access 'kid' use token.Header["kid"]
	// This function is especially useful if you have multiple keys (say for various signing methods - RSA, HMAC, etc...).
	// The standard is to use 'kid' from the token's Header to identify which key to use.
	// However, the parsed token (header and claims) is provided to the callback, thus extending flexibility.
	// JWT Header example:
	// {
	//   "alg":"RS256",
	//   "typ":"JWT",
	//   "kid":"M_GO0JNz4iRvra7NEFI-n"
	// }
	// 'kid' specifies key id/name of the key which should be used for verification.
	// Verification keys (with their name/kid) must be provided by OAuth identity server, which issued the token.

	// What signing method is used in this token?
	if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
		// This method is supported

		// Return RSA Public Key (typically provided to server via config) to be used for JWT verification
		return jwks.GetVerificationPublicKey(token, true)
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		// This method IS NOT SUPPORTED
	}

	if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
		// This method IS NOT SUPPORTED
	}

	return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
}

// parseAndVerifyToken accepts jwt.Claims to fill and parses authorization token string into authorization token struct.
// NB: `claims` is used as an output value
// Provided `claims` is filled with values from the token.
// Token is also verified by the key, provided by getTokenVerificationKey() function
func parseAndVerifyToken(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	// ParseWithClaims takes JWT token (string) and a function which returns key (public) used for JWT verification.
	token, err := jwt.ParseWithClaims(tokenString, claims, getTokenVerificationKey)
	if err != nil {
		log.Errorf("jwt.Parse() FAILED with error %v", err)
		return nil, ErrInvalidToken
	}
	if !token.Valid {
		log.Errorf("jwt.Parse() FAILED with !token.Valid")
		return nil, ErrInvalidToken
	}

	return token, nil
}

/*
func getVerificationCertFromIssuer(token *jwt.Token) (string, error) {
	claims := getCustomizedClaims(token)
	// "https://issuer.url.com/.well-known/jwks.json"
	url := claims.Issuer + "/.well-known/jwks.json"
	resp, err := http.Get(url)
}
*/
