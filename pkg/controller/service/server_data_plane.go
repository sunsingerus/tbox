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
	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/service"
)

// DataPlaneServer specifies default DataPlaneServer
type DataPlaneServer struct {
	service.UnimplementedDataPlaneServer
	// DataChunksHandler is a user-provided handler for DataChunks call
	DataChunksHandler func(service.DataPlane_DataChunksServer, jwt.Claims) error
	// UploadObjectHandler is a user-provided handler for UploadObject call
	UploadObjectHandler func(service.DataPlane_UploadObjectServer, jwt.Claims) error
	// DownloadObjectHandler is a user-provided handler for DownloadObject call
	DownloadObjectHandler func(*common.ObjectRequest, service.DataPlane_DownloadObjectServer, jwt.Claims) error
}

// Verify interface compatibility
var (
	_ service.DataPlaneServer = &DataPlaneServer{}
)

// NewDataPlaneServer creates new DataPlaneServer
func NewDataPlaneServer(
	dataChunksHandler func(service.DataPlane_DataChunksServer, jwt.Claims) error,
	uploadObjectHandler func(service.DataPlane_UploadObjectServer, jwt.Claims) error,
	downloadObjectHandler func(*common.ObjectRequest, service.DataPlane_DownloadObjectServer, jwt.Claims) error,
) *DataPlaneServer {
	return &DataPlaneServer{
		DataChunksHandler:     dataChunksHandler,
		UploadObjectHandler:   uploadObjectHandler,
		DownloadObjectHandler: downloadObjectHandler,
	}
}

// DataChunks gRPC call
func (s *DataPlaneServer) DataChunks(DataChunksServer service.DataPlane_DataChunksServer) error {
	log.Info("DataChunks() - start")
	defer log.Info("DataChunks() - end")

	if s.DataChunksHandler == nil {
		return ErrHandlerUnavailable
	}
	return s.DataChunksHandler(DataChunksServer, ExtractClaims(DataChunksServer.Context()))
}

// UploadObject gRPC call
func (s *DataPlaneServer) UploadObject(UploadObjectServer service.DataPlane_UploadObjectServer) error {
	log.Info("UploadObject() - start")
	defer log.Info("UploadObject() - end")

	if s.UploadObjectHandler == nil {
		return ErrHandlerUnavailable
	}
	return s.UploadObjectHandler(UploadObjectServer, ExtractClaims(UploadObjectServer.Context()))
}

// DownloadObject gRPC call
func (s *DataPlaneServer) DownloadObject(request *common.ObjectRequest, DownloadObjectServer service.DataPlane_DownloadObjectServer) error {
	log.Info("DownloadObject() - start")
	defer log.Info("DownloadObject() - end")

	if s.DownloadObjectHandler == nil {
		return ErrHandlerUnavailable
	}
	return s.DownloadObjectHandler(request, DownloadObjectServer, ExtractClaims(DownloadObjectServer.Context()))
}
