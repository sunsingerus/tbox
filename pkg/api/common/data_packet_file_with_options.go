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
	"bytes"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

// DataPacketFileWithOptions
// Inspired by os.File handler and is expected to be used in the same context.
type DataPacketFileWithOptions struct {
	*DataPacketFile
	Compressor *Compressor
}

// Ensure interface compatibility
var (
	_ IDataChunkFile      = &DataPacketFileWithOptions{}
	_ IDataChunkTransport = &DataPacketFileWithOptions{}
)

// newDataPacketFileWithOptions is unexported constructor
func newDataPacketFileWithOptions(f *DataPacketFile) *DataPacketFileWithOptions {
	return &DataPacketFileWithOptions{
		DataPacketFile: f,
	}
}

// OpenDataPacketFileWOptions opens file with options
// `writer` specifies entity where this file will write data into. Optional, in case the file is read-only
// `reader` specifies entity where this file will read data from. Optional, in case the file is write-only
func OpenDataPacketFileWOptions(
	writer DataPacketWriter,
	reader DataPacketReader,
	options *DataPacketFileOptions,
) (*DataPacketFileWithOptions, error) {
	log.Tracef("OpenDataPacketFileWOptions() - start")
	defer log.Tracef("OpenDataPacketFileWOptions() - end")

	// Open underlying DataPacketFile
	f, err := OpenDataPacketFile(writer, reader)
	if err != nil {
		log.Warnf("FAILED to open DataPacketFile. err: %v", err)
		return nil, err
	}
	f.SetPayloadMetadata(options.GetMetadata())

	this := newDataPacketFileWithOptions(f)

	readCompression := CompressionNone
	if options.GetDecompress() {
		log.Infof("requesting LZMA decompression")
		readCompression = CompressionLZMA
	}

	writeCompression := CompressionNone
	if options.GetCompress() {
		log.Infof("requesting LZMA compression")
		writeCompression = CompressionLZMA
		// Set compression in transport metadata
		this.EnsureStreamOptions().SetCompression(CompressionLZMA)
	}
	this.Compressor, err = NewCompressor(
		readCompression,
		this.GetDataPacketFile(),
		writeCompression,
		this.GetDataPacketFile(),
	)
	if err != nil {
		log.Errorf("UNABLE to setup compression options. Err: %v", err)
		return nil, err
	}

	return this, nil
}

// Close is an io.Closer interface function
func (f *DataPacketFileWithOptions) Close() error {
	log.Tracef("DataPacketFileWithOptions.Close() - start")
	defer log.Tracef("DataPacketFileWithOptions.Close() - end")

	if f == nil {
		return nil
	}

	err1 := f.Compressor.Close()
	// Need to explicitly call Close(), because Compressor.Close() does not call Close() on underlying io.Writer()
	err2 := f.GetDataPacketFile().Close()

	switch {
	case err1 != nil:
		return err1
	case err2 != nil:
		return err2
	default:
		return nil
	}
}

// GetDataPacketFile is a getter
func (f *DataPacketFileWithOptions) GetDataPacketFile() *DataPacketFile {
	if f == nil {
		return nil
	}
	return f.DataPacketFile
}

// Write is an io.Writer interface function
func (f *DataPacketFileWithOptions) Write(p []byte) (n int, err error) {
	log.Tracef("DataPacketFileWithOptions.Write() - start: %d", len(p))
	defer log.Tracef("DataPacketFileWithOptions.Write() - end  : %d", len(p))

	if f.Compressor.WriteEnabled() {
		return f.Compressor.Write(p)
	}
	if f.GetDataPacketFile() != nil {
		return f.GetDataPacketFile().Write(p)
	}

	return 0, fmt.Errorf("unknown write() entity")
}

// WriteTo is an io.WriterTo interface function
func (f *DataPacketFileWithOptions) WriteTo(dst io.Writer) (int64, error) {
	log.Tracef("DataPacketFileWithOptions.WriteTo() - start")
	defer log.Tracef("DataPacketFileWithOptions.WriteTo() - end")

	return cp(dst, f)
}

// Read is an io.Reader interface function
func (f *DataPacketFileWithOptions) Read(p []byte) (n int, err error) {
	log.Tracef("DataPacketFileWithOptions.Read() - start")
	defer log.Tracef("DataPacketFileWithOptions.Read() - end")

	if f.Compressor.ReadEnabled() {
		// TODO need to read uncompressed data even if compression requested
		log.Debugf("decompression requested")

		if f.GetStreamOptions() == nil {
			log.Debugf("no stream options yet, wait for it")
			f.receiveIDataChunkIntoBuf()
		}

		if f.GetStreamOptions() == nil {
			log.Warnf("got no stream options, abort")
			return 0, fmt.Errorf("decompression requested, but no metadata available")
		}

		if f.GetStreamOptions().HasCompression() {
			log.Tracef("reading compressed data")
			return f.Compressor.Read(p)
		}

		log.Warnf("unknown compression method %v", f.GetStreamOptions().GetCompression().GetType())

		return 0, fmt.Errorf("unknown compression method")
	}

	if f.GetDataPacketFile() != nil {
		return f.GetDataPacketFile().Read(p)
	}

	return 0, fmt.Errorf("unknown read() entity")
}

// ReadFrom is an io.ReaderFrom interface function
func (f *DataPacketFileWithOptions) ReadFrom(src io.Reader) (int64, error) {
	log.Tracef("DataPacketFileWithOptions.ReadFrom() - start")
	defer log.Tracef("DataPacketFileWithOptions.ReadFrom() - end")

	n, err := cp(f, src)

	log.Debugf("Accepted data meta:")
	f.GetPayloadMetadata().Log()
	f.GetStreamOptions().Log()

	return n, err
}

// WriteToBuf writes data to newly created buffer
func (f *DataPacketFileWithOptions) WriteToBuf() (int64, *bytes.Buffer, error) {
	log.Tracef("DataPacketFileWithOptions.WriteToBuf() - start")
	defer log.Tracef("DataPacketFileWithOptions.WriteToBuf() - end")

	var buf = &bytes.Buffer{}
	written, err := f.WriteTo(buf)
	if err != nil {
		log.Errorf("got error: %v", err.Error())
	}

	// Debug
	log.Debugf("metadata: %s", f.GetPayloadMetadata())
	log.Debugf("data: %s", buf.String())

	return written, buf, err
}
