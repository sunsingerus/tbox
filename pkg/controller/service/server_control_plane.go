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
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/api/service"
	"github.com/sunsingerus/tbox/pkg/controller"
)

func GetOutgoingQueue() chan *common.Task {
	return controller.GetOutgoing()
}

func GetIncomingQueue() chan *common.Task {
	return controller.GetIncoming()
}

// ControlPlaneServer specifies default ControlPlaneServer
type ControlPlaneServer struct {
	service.UnimplementedControlPlaneServer
}

// NewControlPlaneServer creates new ControlPlaneServer
func NewControlPlaneServer() *ControlPlaneServer {
	return &ControlPlaneServer{}
}

// TasksHandler is a user-provided handler for Tasks call
var TasksHandler = func(service.ControlPlane_TasksServer, jwt.Claims) error {
	return ErrHandlerUnavailable
}

// Tasks gRPC call
func (s *ControlPlaneServer) Tasks(TasksServer service.ControlPlane_TasksServer) error {
	log.Info("Tasks() - start")
	defer log.Info("Tasks() - end")

	if TasksHandler == nil {
		return ErrHandlerUnavailable
	}
	return TasksHandler(TasksServer, ExtractClaims(TasksServer.Context()))

	// controller.CommandsExchangeEndlessLoop(CommandsServer)
	// return nil
}
