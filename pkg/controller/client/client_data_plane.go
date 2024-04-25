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
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/api/service"
)

// DataExchange send data to server and receives back reply (if needed)
func DataExchange(
	DataPlaneClient service.DataPlaneClient,
	src io.Reader,
	options *DataExchangeOptions,
) *DataExchangeResult {
	log.Infof("DataExchange() - start")
	defer log.Infof("DataExchange() - end")

	result := NewDataExchangeResult()

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var DataChunksBiMultiClient service.DataPlane_DataChunksClient

	DataChunksBiMultiClient, result.Error = DataPlaneClient.DataChunks(ctx)
	if result.Error != nil {
		log.Errorf("ControlPlaneClient.DataChunks() failed %v", result.Error)
		return result
	}

	defer func() {
		// This is hand-made flush() replacement for gRPC
		// It is required in order to flush all outstanding data before
		// context's cancel() is called, which simply discards all outstanding data.
		// On receiving end, when cancel() is the first in the race, f receives 'cancel' and (sometimes) no data
		// instead of complete set of data and EOF
		// See https://github.com/grpc/grpc-go/issues/1714 for more details
		DataChunksBiMultiClient.CloseSend()
		DataChunksBiMultiClient.Recv()
	}()

	f, err := common.OpenDataPacketFileWOptions(
		DataChunksBiMultiClient,
		DataChunksBiMultiClient,
		common.NewDataPacketFileOptions().
			SetMetadata(options.GetMetadata()).
			SetCompress(options.GetCompress()).
			SetDecompress(options.GetDecompress()),
	)
	if err != nil {
		log.Errorf("DataPlaneClient.DataExchange() failed %v", result.Error)
		result.Error = err
		return result
	}

	if src != nil {
		// We have something to send
		result.Send.Data.Len, result.Error = f.ReadFrom(src)
		if result.Error != nil {
			log.Warnf("DataPlaneClient.DataExchange() failed with err %v", result.Error)
			return result
		}
	}

	if options.GetWaitReply() {
		// We should wait for reply
		result.Recv.Data.Len, result.Recv.Data.Data, result.Error = f.WriteToBuf()
		if result.Error != nil {
			log.Warnf("DataPlaneClient.DataExchange() failed with err %v", result.Error)
			return result
		}
	}
	result.Recv.Data.Metadata = f.GetPayloadMetadata()

	return result
}

// Upload send data to server and receives back status(es)
func Upload(
	DataPlaneClient service.DataPlaneClient,
	src io.Reader,
	options *DataExchangeOptions,
) *DataExchangeResult {
	log.Infof("Upload() - start")
	defer log.Infof("Upload() - end")

	result := NewDataExchangeResult()

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var DataChunksUpOneClient service.DataPlane_UploadObjectClient
	DataChunksUpOneClient, result.Error = DataPlaneClient.UploadObject(ctx)
	if result.Error != nil {
		log.Errorf("DataPlaneClient.Upload() failed %v", result.Error)
		return result
	}

	f, err := common.OpenDataPacketFileWOptions(
		DataChunksUpOneClient,
		nil,
		common.NewDataPacketFileOptions().
			SetMetadata(options.GetMetadata()).
			SetCompress(options.GetCompress()).
			SetDecompress(options.GetDecompress()),
	)
	if err != nil {
		log.Errorf("DataPlaneClient.Upload() failed %v", result.Error)
		result.Error = err
		return result
	}

	if src != nil {
		// We have something to send
		result.Send.Data.Len, result.Error = f.ReadFrom(src)
		if result.Error != nil {
			log.Warnf("DataPlaneClient.Upload() failed with err %v", result.Error)
			return result
		}
	}

	result.Recv.ObjectStatus, result.Error = DataChunksUpOneClient.CloseAndRecv()

	return result
}

// Download downloads data from server
func Download(
	DataPlaneClient service.DataPlaneClient,
	dst io.Writer,
	taskId, filename string,
	options *DataExchangeOptions,
) *DataExchangeResult {
	log.Infof("Download() - start")
	defer log.Infof("Download() - end")

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	request := common.NewObjectRequest().
		AppendAddress(
			common.DomainTaskID,
			common.NewAddress().Set(common.NewUuidFromString(taskId)),
		).
		AppendAddress(
			common.DomainFilename,
			common.NewAddress().Set(common.NewFilename(filename)),
		)

	var client service.DataPlane_DownloadObjectClient
	result := NewDataExchangeResult()

	client, result.Error = DataPlaneClient.DownloadObject(ctx, request)
	if result.Error != nil {
		log.Errorf("DataPlaneClient.Download() failed %v", result.Error)
		return result
	}

	f, err := common.OpenDataPacketFileWOptions(
		nil,
		client,
		common.NewDataPacketFileOptions().
			SetMetadata(options.GetMetadata()).
			SetCompress(options.GetCompress()).
			SetDecompress(options.GetDecompress()),
	)
	if err != nil {
		log.Errorf("DataPlaneClient.Download() failed %v", result.Error)
		result.Error = err
		return result
	}

	if dst != nil {
		// We have something to send
		result.Recv.Data.Len, result.Error = f.WriteTo(dst)
		if result.Error != nil {
			log.Warnf("DataPlaneClient.Download() failed with err %v", result.Error)
			return result
		}
		result.Recv.Data.Metadata = f.GetPayloadMetadata()
	}

	return result
}
