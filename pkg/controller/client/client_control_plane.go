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

package controller_client

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/service"
	"github.com/sunsingerus/tbox/pkg/controller"
)

// TasksExchange exchanges tasks
func TasksExchange(ControlPlaneClient service.ControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//this code sends token per each RPC call:
	//	md := metadata.Pairs("authorization", "my-secret-token")
	//	ctx = metadata.NewOutgoingContext(ctx, md)

	rpcTasks, err := ControlPlaneClient.Tasks(ctx)
	if err != nil {
		log.Fatalf("ControlPlaneClient.Tasks() failed %v", err)
		os.Exit(1)
	}
	defer rpcTasks.CloseSend()

	log.Infof("Tasks() called")
	controller.TasksExchangeEndlessLoop(rpcTasks)
}
