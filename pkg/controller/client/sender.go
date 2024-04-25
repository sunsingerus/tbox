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
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/api/service"
)

// SendFile sends file from client to service and receives response back (if any)
func SendFile(client service.DataPlaneClient, filename string, options *DataExchangeOptions) (int64, error) {
	log.Info("SendFile() - start")
	defer log.Info("SendFile() - end")

	if _, err := os.Stat(filename); err != nil {
		log.Warnf("no file %s available err: %v", filename, err)
		return 0, err
	}

	log.Infof("Has file %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Warnf("ERROR open file %s err: %v", filename, err)
		return 0, err
	}

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(filepath.Base(filename))
	return SendReader(client, f, options)
}

// SendStdin sends STDIN from client to service and receives response back (if any)
func SendStdin(client service.DataPlaneClient, options *DataExchangeOptions) (int64, error) {
	log.Info("SendStdin() - start")
	defer log.Info("SendStdin() - end")

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(os.Stdin.Name())
	return SendReader(client, os.Stdin, options)
}

// SendReader
func SendReader(client service.DataPlaneClient, r io.Reader, options *DataExchangeOptions) (int64, error) {
	log.Info("SendReader() - start")
	defer log.Info("SendReader() - end")

	result := DataExchange(client, r, options)
	if result.Error == nil {
		log.Infof("DONE send %s size %d", "io.Reader", result.Send.Data.Len)
	} else {
		log.Warnf("FAILED send %s size %d err %v", "io.Reader", result.Send.Data.Len, result.Error)
	}

	return result.Send.Data.Len, result.Error
}

// SendBytes
func SendBytes(client service.DataPlaneClient, data []byte, options *DataExchangeOptions) (int64, error) {
	log.Info("SendBytes() - start")
	defer log.Info("SendBytes() - end")

	return SendReader(client, bytes.NewReader(data), options)
}

// SendEchoRequest
func SendEchoRequest(outgoingQueue chan *common.Task) {
	for i := 0; i < 5; i++ {
		task := common.NewTask().SetType(common.TaskEchoRequest).CreateUuid().SetDescription("desc")
		outgoingQueue <- task

		log.Infof("Wait before send new Echo Request")
		time.Sleep(3 * time.Second)
	}
}
