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
	"github.com/sunsingerus/tbox/pkg/api/service"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

// DownloadIntoFile sends file from client to service and receives response back (if any)
func DownloadIntoFile(client service.DataPlaneClient, taskId, filename string, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("DownloadIntoFile() - start")
	defer log.Info("DownloadIntoFile() - end")

	f, err := os.Open(filename)
	if err != nil {
		log.Warnf("ERROR open file %s err: %v", filename, err)
		return NewDataExchangeResultError(err)
	}
	defer f.Close()

	return DownloadIntoWriter(client, taskId, filename, f, options)
}

// DownloadIntoStdout sends STDIN from client to service and receives response back (if any)
func DownloadIntoStdout(client service.DataPlaneClient, taskId, filename string, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("DownloadIntoStdout() - start")
	defer log.Info("DownloadIntoStdout() - end")

	return DownloadIntoWriter(client, taskId, filename, os.Stdout, options)
}

// DownloadIntoBuffer downloads into bytes.Buffer of DataExchangeResult
func DownloadIntoBuffer(client service.DataPlaneClient, taskId, filename string, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("DownloadIntoBuffer() - start")
	defer log.Info("DownloadIntoBuffer() - end")

	buf := &bytes.Buffer{}

	res := DownloadIntoWriter(client, taskId, filename, buf, options)
	res.Recv.Data.Len = int64(buf.Len())
	res.Recv.Data.Data = buf
	return res
}

// DownloadIntoWriter downloads into io.Writer
func DownloadIntoWriter(client service.DataPlaneClient, taskId, filename string, w io.Writer, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("DownloadIntoWriter() - start")
	defer log.Info("DownloadIntoWriter() - end")

	result := Download(client, w, taskId, filename, options)
	if result.Error == nil {
		log.Infof("DONE send %s size %d", "io.Reader", result.Send.Data.Len)
	} else {
		log.Warnf("FAILED send %s size %d err %v", "io.Reader", result.Send.Data.Len, result.Error)
	}

	return result
}
