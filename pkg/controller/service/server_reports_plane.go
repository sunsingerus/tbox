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

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/api/service"
)

// ReportsPlaneServer specifies default ReportsPlaneServer
type ReportsPlaneServer struct {
	service.UnimplementedReportsPlaneServer
}

// NewReportsPlaneServer creates new ReportsPlaneServer
func NewReportsPlaneServer() *ReportsPlaneServer {
	return &ReportsPlaneServer{}
}

// ObjectsReportHandler is a user-provided handler for ObjectsReport call
var ObjectsReportHandler = func(context.Context, *common.ObjectsRequest, jwt.Claims) (*common.ObjectsList, error) {
	return nil, ErrHandlerUnavailable
}

// ObjectsReport gRPC call
func (s *ReportsPlaneServer) ObjectsReport(ctx context.Context, req *common.ObjectsRequest) (*common.ObjectsList, error) {
	log.Info("ObjectsReport() - start")
	defer log.Info("ObjectsReport() - end")

	if ObjectsReportHandler == nil {
		return nil, ErrHandlerUnavailable
	}
	return ObjectsReportHandler(ctx, req, ExtractClaims(ctx))
}
