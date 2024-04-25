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

package common

import (
	"fmt"
	"io"
)

// IDataChunk represents DataChunk interface
type IDataChunk interface {
	// GetDataLen gets length of the data of the DataChunk
	GetDataLen() int
	// GetData gets data of the DataChunk
	GetData() []byte
	// SetData sets data of the DataChunk
	SetData([]byte)
	// GetLast returns whether this DataChunk is the last in the file (set)
	GetLast() bool
	// SetLast sets flag whether this DataChunk is the last in the file (set)
	SetLast(bool)
	// GetOffset gets offset of the DataChunk within the file (set)
	GetOffset() int64
	// SetOffset sets offset of the DataChunk within the file (set)
	SetOffset(int64)
	// Stringer is a nice touch to log DataChunk(s)
	fmt.Stringer
}

// IDataChunkFile represents generic DataChunkFile which can transport IDataChunk(s)
type IDataChunkFile interface {
	GetOffsetter

	// Streaming interfaces allows DataChunkFiler to be used seamlessly with file/stream-like functions.

	io.Writer
	io.WriterTo
	io.Reader
	io.ReaderFrom
	io.Closer
}

type GetOffsetter interface {
	// GetOffset gets current offset within the file (set)
	GetOffset() int64
}

// IDataChunkTransport represents interface to transport abstracted data via IDataChunk.
// Interfaces transport functions to Receive, Send and create abstracted IDataChunk.
type IDataChunkTransport interface {
	// NewIDataChunk creates new adbstracted data chunk. Typically, to be sent at the transport level by Send()
	NewIDataChunk(GetOffsetter) IDataChunk
	// Recv receives incoming abstracted IDataChunk at the transport layer.
	Recv() (IDataChunk, error)
	// Send sends IDataChunk at the transport layer.
	Send(IDataChunk) error
}
