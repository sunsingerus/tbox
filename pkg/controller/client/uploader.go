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
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/service"
)

// UploadFile sends file from client to service and receives response back (if any)
func UploadFile(client service.DataPlaneClient, filename string, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadFile() - start")
	defer log.Info("UploadFile() - end")

	if _, err := os.Stat(filename); err != nil {
		log.Warnf("no file %s available err: %v", filename, err)
		return NewDataExchangeResultError(err)
	}

	log.Infof("Has file %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Warnf("ERROR open file %s err: %v", filename, err)
		return NewDataExchangeResultError(err)
	}
	defer f.Close()

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(filepath.Base(filename))
	return UploadReader(client, f, options)
}

// UploadStdin sends STDIN from client to service and receives response back (if any)
func UploadStdin(client service.DataPlaneClient, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadStdin() - start")
	defer log.Info("UploadStdin() - end")

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(os.Stdin.Name())
	return UploadReader(client, os.Stdin, options)
}

// UploadStdinLog sends STDIN in time chunks from client to service and receives response back (if any)
func UploadStdinLog(client service.DataPlaneClient, options *DataExchangeOptions) *DataExchangeResult {
	type ev struct {
		str string
		eof bool
	}

	log.Info("UploadStdinLog() - start")
	defer log.Info("UploadStdinLog() - end")

	options = options.Ensure()
	options.EnsureMetadata().SetFilename(os.Stdin.Name())

	stdin := make(chan ev) // channel with stdin strings
	go func() {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			stdin <- ev{scan.Text(), false}
		}
		stdin <- ev{"", true}
		close(stdin)
		log.Info("UploadStdinLog() - goroutine finished")
	}()

	toSend := bytes.NewBuffer(nil)
	timer := time.NewTimer(3 * time.Second)
	ticker := time.NewTicker(30 * time.Second)
	sizeSignal := make(chan struct{})

MAIN:
	for {
		select {

		case s := <-stdin:
			if s.eof {
				break MAIN
			}
			toSend.WriteString(s.str + "\n")
			if toSend.Len() > 8192 {
				go func() { sizeSignal <- struct{}{} }()
			}

		case <-timer.C:
			if toSend.Len() > 0 {
				_ = UploadReader(client, toSend, options)
				toSend = bytes.NewBuffer(nil)
			}

		case <-ticker.C:
			if toSend.Len() > 0 {
				_ = UploadReader(client, toSend, options)
				toSend = bytes.NewBuffer(nil)
			}

		case <-sizeSignal:
			_ = UploadReader(client, toSend, options)
			toSend = bytes.NewBuffer(nil)

		}

		timer = time.NewTimer(3 * time.Second)
	}

	if toSend.Len() == 0 {
		return &DataExchangeResult{}
	}

	return UploadReader(client, toSend, options)
}

// UploadBytes
func UploadBytes(client service.DataPlaneClient, data []byte, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadBytes() - start")
	defer log.Info("UploadBytes() - end")

	return UploadReader(client, bytes.NewReader(data), options)
}

// UploadReader
func UploadReader(client service.DataPlaneClient, r io.Reader, options *DataExchangeOptions) *DataExchangeResult {
	log.Info("UploadReader() - start")
	defer log.Info("UploadReader() - end")

	result := Upload(client, r, options)
	if result.Error == nil {
		log.Infof("DONE send %s size %d", "io.Reader", result.Send.Data.Len)
	} else {
		log.Warnf("FAILED send %s size %d err %v", "io.Reader", result.Send.Data.Len, result.Error)
	}

	return result
}
