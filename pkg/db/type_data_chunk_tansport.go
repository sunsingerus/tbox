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

package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
)

// DataChunkTransport defines transport level interface
// Has the following functions:
//   Send(*DataChunk) error
//   Recv() (*DataChunk, error)

type DataChunkTransport struct {
}

// NewDataChunkTransport
func NewDataChunkTransport() *DataChunkTransport {
	log.Infof("clickhouse.NewDataChunkTransport() - start")
	defer log.Infof("clickhouse.NewDataChunkTransport() - end")

	return &DataChunkTransport{}
}

// Send
func (t *DataChunkTransport) Send(dataChunk *common.DataChunk) error {
	log.Infof("clickhouse.DataChunkTransport.Send() - start")
	defer log.Infof("clickhouse.DataChunkTransport.Send() - end")

	return nil
}

// Recv
func (t *DataChunkTransport) Recv() (*common.DataChunk, error) {
	log.Infof("clickhouse.DataChunkTransport.Recv() - start")
	defer log.Infof("clickhouse.DataChunkTransport.Recv() - end")

	return &common.DataChunk{}, nil
}
