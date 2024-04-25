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
	"github.com/sunsingerus/tbox/pkg/util"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
)

const defaultMaxWriteIDataChunkSize = 1024

// DataChunkFile is a generic (abstract) file (set) of IDataChunk(s)
// Inspired by os.File handler and is expected to be used in the same context.
// Implements the following interfaces:
//   - io.Writer
//   - io.WriterTo
//   - io.Reader
//   - io.ReaderFrom
//   - io.Closer
//
// and thus can be used in any functions, which operate these interfaces, such as io.Copy()
// It is a base for creating handlers for data files (sets, streams)
type DataChunkFile struct {
	// transport provides functions to transport IDataChunk(s)
	transport IDataChunkTransport

	// globalInitialOffset of this set of IDataChunk(s) within something bigger. Optional.
	globalInitialOffset int64
	// globalOffset of the current IDataChunk(s) within something bigger. Optional.
	// globalOffset = globalInitialOffset + offset
	globalOffset int64
	// offset of the current IDataChunk within current set of IDataChunk(s)
	offset int64

	// maxWriteIDataChunkSize limits max size of a payload within one data IDataChunk to be sent
	maxWriteIDataChunkSize int

	//
	// Receiver section
	//

	// receivedDataBuf is a buf, where data of incoming IDataChunk(s) is accumulated.
	// Data is being read by caller(s) from this buf.
	receivedDataBuf []byte
	// receivedDataErr is an error occurred upon receiving data
	receivedDataErr error

	//
	// Log section
	//

	// lastTimeDataIDataChunkLogged specifies when last time data chunk was logged
	lastTimeDataIDataChunkLogged time.Time
}

// Ensure interface compatibility
var (
	_ IDataChunkFile = &DataChunkFile{}
)

// OpenDataChunkFile opens DataChunk(s) file
// Inspired by os.OpenFile()
func OpenDataChunkFile(transport IDataChunkTransport) (*DataChunkFile, error) {
	return &DataChunkFile{
		transport: transport,
	}, nil
}

// GetOffset current offset
func (f *DataChunkFile) GetOffset() int64 {
	if f == nil {
		return 0
	}
	return f.offset
}

// close does internal job to close the communication
func (f *DataChunkFile) close() {
	if f == nil {
		return
	}

	f.globalInitialOffset = 0
	f.globalOffset = 0
	f.offset = 0

	f.receivedDataBuf = nil
	f.receivedDataErr = nil

	// TODO flush outgoing transport
	// TODO should incoming transport be drained?
}

// logIDataChunk logs IDataChunk. Aux functionality.
func (f *DataChunkFile) logIDataChunk(chunk IDataChunk) {
	if chunk == nil {
		return
	}

	now := time.Now()
	interval := 30 * time.Second
	switch {
	case
		f.lastTimeDataIDataChunkLogged.IsZero(),                 // No chunks logged before, log the first one
		chunk.GetLast(),                                         // Log the last one to complete transfer logging
		now.After(f.lastTimeDataIDataChunkLogged.Add(interval)), // Log every X seconds
		log.GetLevel() == log.TraceLevel:                        // In case of trace all chunks should be logged
		// Should log this DataChunk
	default:
		// Should not log this DataChunk
		return
	}

	log.Infof("got IDataChunk. %s", chunk)
	f.lastTimeDataIDataChunkLogged = now
}

// receiveIDataChunk waits for one IDataChunk and reads incoming IDataChunk.
// Is a blocking function.
func (f *DataChunkFile) receiveIDataChunk() (IDataChunk, error) {
	log.Tracef("DataChunkFile.receiveIDataChunk() - start")
	defer log.Tracef("DataChunkFile.receiveIDataChunk() - end")

	// Receive abstracted IDataChunk from the reader
	iDataChunk, err := f.transport.Recv()
	f.logIDataChunk(iDataChunk)

	switch err {
	case nil:
		// All went well, ready to receive more data
	case io.EOF:
		// Correct EOF arrived
		if iDataChunk.GetDataLen() == 0 {
			log.Infof("DataChunkFile.receiveIDataChunk() get EOF with no data")
		} else {
			log.Infof("DataChunkFile.receiveIDataChunk() get EOF with %d bytes", iDataChunk.GetDataLen())
		}
	default:
		// Stream is somehow broken
		log.Infof("DataChunkFile.receiveIDataChunk() got err: %v", err)
	}

	return iDataChunk, err
}

// receiveIDataChunkIntoBuf is used to wait for and to read into buf data chunk.
func (f *DataChunkFile) receiveIDataChunkIntoBuf() {
	log.Tracef("DataChunkFile.receiveIDataChunkIntoBuf() - start")
	defer log.Tracef("DataChunkFile.receiveIDataChunkIntoBuf() - end")

	// Wait for a DataChunk to arrive
	iDataChunk, err := f.receiveIDataChunk()
	if iDataChunk != nil {
		log.Tracef("Got data len: %d", iDataChunk.GetDataLen())
		if iDataChunk.GetDataLen() > 0 {
			f.receivedDataBuf = append(f.receivedDataBuf, iDataChunk.GetData()...)
		}

		if iDataChunk.GetLast() {
			log.Tracef("Got last IDataChunk, reporting EOF ")
			f.receivedDataErr = io.EOF
		}
	}

	if err != nil {
		log.Tracef("Got an err: %v", err)
		f.receivedDataErr = err
	}
}

// sendIDataChunk is used to send data buf as one IDataChunk.
// Returns an error, in case provided data buf to send is bigger than one IDataChunk accepts.
func (f *DataChunkFile) sendIDataChunk(p []byte) (n int, err error) {
	log.Tracef("DataChunkFile.sendIDataChunk() - start")
	defer log.Tracef("DataChunkFile.sendIDataChunk() - end")

	n = len(p)

	// Some sanity checks
	if n == 0 {
		return 0, nil
	}
	if (n > f.maxWriteIDataChunkSize) && (f.maxWriteIDataChunkSize > 0) {
		return 0, fmt.Errorf("attempt to sendIDataChunk() with oversized chunk: %d > %d", n, f.maxWriteIDataChunkSize)
	}

	// Offset of this data chunk
	f.globalOffset = f.globalInitialOffset + f.offset

	// Create abstracted IDataChunk
	iDataChunk := f.transport.NewIDataChunk(f)
	iDataChunk.SetOffset(f.offset)
	iDataChunk.SetData(p)
	err = f.transport.Send(iDataChunk)
	if err != nil {
		// We have some kind of error and were not able to send the data
		n = 0
		if err == io.EOF {
			// Not sure, probably this is not an error, after all
			log.Infof("sendIDataChunk.Send() received EOF")
		} else {
			log.Errorf("sendIDataChunk.Send() FAILED %v", err)
		}
	}

	// Adjust offset of next chunk to be sent after this one
	f.offset += int64(n)
	return
}

// iDataChunkSize calculates size for data buf of length bufLen
func (f *DataChunkFile) iDataChunkSize(bufLen int) int {
	if f.maxWriteIDataChunkSize <= 0 {
		// MaxWriteIDataChunkSize is not specified, return the whole buf length
		return bufLen
	}

	// Max chunk size is specified, result size must not be greater than MaxWriteIDataChunkSize
	return util.Min(f.maxWriteIDataChunkSize, bufLen)
}

// Write implements io.Writer
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write process to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the data slice, even temporarily.
//
// Implementations must not retain p.
func (f *DataChunkFile) Write(p []byte) (n int, err error) {
	log.Tracef("DataChunkFile.Write() - start")
	defer log.Tracef("DataChunkFile.Write() - end")

	// Keep sending data while we have
	for len(p) > 0 {
		if sent, e := f.sendIDataChunk(p[:f.iDataChunkSize(len(p))]); e == nil {
			// Chunk sent, slide the window
			n += sent
			p = p[sent:]
		} else {
			// Unable to send
			return n, e
		}
	}

	return n, nil
}

// WriteTo implements io.WriterTo
//
// io.WriterTo is the interface that wraps the WriteTo method.
//
// WriteTo writes data to dst until there's no more data to write or
// when an error occurs. The return value n is the number of bytes
// written. Any error encountered during the write is also returned.
//
// The Copy function uses WriterTo if available.
func (f *DataChunkFile) WriteTo(dst io.Writer) (n int64, err error) {
	log.Tracef("DataChunkFile.WriteTo() - start")
	defer log.Tracef("DataChunkFile.WriteTo() - end")

	return cp(dst, f, util.IfPositiveValue(f.maxWriteIDataChunkSize, defaultMaxWriteIDataChunkSize))
}

// Read implements io.Reader
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.
func (f *DataChunkFile) Read(p []byte) (n int, err error) {
	log.Tracef("DataChunkFile.Read() - start")
	defer log.Tracef("DataChunkFile.Read() - end")

	if len(f.receivedDataBuf) == 0 {
		// No buffered data available, need to get some
		f.receiveIDataChunkIntoBuf()
	}

	n = 0
	if len(f.receivedDataBuf) > 0 {
		// Have some buffered data, copy it out
		n = copy(p, f.receivedDataBuf)
		// Cut off data that were just copied out
		f.receivedDataBuf = f.receivedDataBuf[n:]
	}

	if len(f.receivedDataBuf) > 0 {
		// Still have some data ready for next read, so not even EOF should be reported right now
		return n, nil
	}

	// No more data available for immediate read
	return n, f.receivedDataErr
}

// ReadFrom implements io.ReaderFrom
//
// io.ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from src until EOF or error.
// The return value n is the number of bytes read.
// Any error except io.EOF encountered during the read is also returned.
//
// The Copy function uses ReaderFrom if available.
func (f *DataChunkFile) ReadFrom(src io.Reader) (n int64, err error) {
	log.Tracef("DataChunkFile.ReadFrom() - start")
	defer log.Tracef("DataChunkFile.ReadFrom() - end")

	return cp(f, src, util.IfPositiveValue(f.maxWriteIDataChunkSize, defaultMaxWriteIDataChunkSize))
}

// Close implements io.Closer
//
// io.Closer is the interface that wraps the basic Close method.
//
// The behavior of Close after the first call is undefined.
// Specific implementations may document their own behavior.
func (f *DataChunkFile) Close() error {
	log.Tracef("DataChunkFile.Close() - start")
	defer log.Tracef("DataChunkFile.Close() - end")
	if f == nil {
		return nil
	}

	defer f.close()

	if f.offset == 0 {
		// No data were sent via this stream, no need to send finalizer
		return nil
	}

	// Some data were sent via this stream, need to finalize transmission with finalizer

	// Send "last" data chunk
	// Create abstracted IDataChunk
	iDataChunk := f.transport.NewIDataChunk(f)
	iDataChunk.SetLast(true)
	err := f.transport.Send(iDataChunk)
	if err != nil {
		if err == io.EOF {
			log.Infof("Send() received EOF")
		} else {
			log.Errorf("failed to Send() %v", err)
		}
	}

	return err
}

// TODO implement Reset function for DataChunkFile,
//  so the same descriptor can be used for multiple transmissions.
