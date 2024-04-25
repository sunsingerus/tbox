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

package controller

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
)

// TaskSenderReceiver defines transport level interface (for both client and server),
// which serves Tasks streams bi-directionally.
type TaskSenderReceiver interface {
	Send(*common.Task) error
	Recv() (*common.Task, error)
}

func TasksExchangeEndlessLoop(TaskSenderReceiver TaskSenderReceiver) {
	waitIncoming := make(chan bool)
	waitOutgoing := make(chan bool)

	// Recv() loop
	go func() {
		for {
			msg, err := TaskSenderReceiver.Recv()
			if msg != nil {
				log.Infof("TasksExchangeEndlessLoop.Recv() got msg")
				GetIncoming() <- msg
			}
			if err == nil {
				// All went well, ready to receive more data
			} else if err == io.EOF {
				// Correct EOF
				log.Infof("TasksExchangeEndlessLoop.Recv() got EOF")

				close(waitIncoming)
				return
			} else {
				// Stream broken
				log.Infof("TasksExchangeEndlessLoop.Recv() got err: %v", err)

				close(waitIncoming)
				return
			}
		}
	}()

	// Send() loop
	go func() {
		for {
			select {
			case <-waitIncoming:
				// Incoming stream from this client is closed/broken, no need to wait commands for it
				close(waitOutgoing)
				return
			case task := <-GetOutgoing():
				log.Infof("got task to send")
				err := TaskSenderReceiver.Send(task)
				if err == nil {
					// All went well
					log.Infof("TasksExchangeEndlessLoop.Send() OK")
				} else if err == io.EOF {
					log.Infof("TasksExchangeEndlessLoop.Send() got EOF")

					close(waitOutgoing)
					return
				} else {
					log.Fatalf("TasksExchangeEndlessLoop.Send() got err: %v", err)

					close(waitOutgoing)
					return
				}
			}
		}
	}()

	<-waitIncoming
	<-waitOutgoing
}

var (
	incomingBacklog int32 = 100
	incoming        chan *common.Task
	outgoingBacklog int32 = 100
	outgoing        chan *common.Task
)

func Init() {
	incoming = make(chan *common.Task, incomingBacklog)
	outgoing = make(chan *common.Task, outgoingBacklog)
}

func GetOutgoing() chan *common.Task {
	return outgoing
}

func GetIncoming() chan *common.Task {
	return incoming
}
