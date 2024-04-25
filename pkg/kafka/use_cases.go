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

package kafka

import (
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
)

// CopyDataChunkFile
func CopyDataChunkFile(consumer *Consumer, dst io.Writer) {
	transport := NewDataChunkTransport(
		nil,
		consumer,
		true,
	)
	defer transport.Close()

	f, err := common.OpenDataPacketFile(transport, transport)
	if err != nil {
		log.Errorf("err: %v", err)
	}
	defer f.Close()

	n, err := io.Copy(dst, f)
	if err == nil {
		log.Infof("written: %d", n)
		f.GetPayloadMetadata().Log()
	} else {
		log.Errorf("err: %v", err)
	}
}

func TasksProcessor(consumer *Consumer, processor func(*common.Task) error) {
	transport := NewTaskTransport(
		nil,
		consumer,
		true,
	)
	defer transport.Close()

	for {
		task, err := transport.Recv()
		if task != nil {
			processor(task)
		}
		if err != nil {
			return
		}
	}
}
