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

	log "github.com/sirupsen/logrus"
	"github.com/ulikunitz/xz/lzma"
)

// Compressor is a compression descriptor
type Compressor struct {
	ReadCompression  *Compression
	LZMAReader       *lzma.Reader
	WriteCompression *Compression
	LZMAWriter       *lzma.Writer
}

// NewCompressor creates new compressor(s) for provided io.Reader and io.Writer
// When data are read from compressor by calling Read() of the compressor,
// compressor reads compressed data from `reader`, inflates it and returns as the result of its (compressor's) Read()
// When data are written into compressor by calling Write() of the compressor,
// compressor deflates data and writes compressed data into `writer`
func NewCompressor(
	readCompression *Compression,
	reader io.Reader,
	writeCompression *Compression,
	writer io.Writer,
) (*Compressor, error) {
	compressor := &Compressor{}

	switch readCompression.GetType() {
	case CompressionTypeLZMA:
		compressor.ReadCompression = readCompression
		lzmaReader, err := lzma.NewReader(reader)
		if err != nil {
			log.Warnf("FAILED to create lzma reader. err: %v", err)
			return nil, err
		}
		compressor.LZMAReader = lzmaReader
	default:
		compressor.ReadCompression = nil
	}

	switch writeCompression.GetType() {
	case CompressionTypeLZMA:
		lzmaWriter, err := lzma.NewWriter(writer)
		if err != nil {
			log.Warnf("FAILED to create lzma writer. err: %v", err)
			return nil, err
		}
		compressor.WriteCompression = writeCompression
		compressor.LZMAWriter = lzmaWriter
	default:
		compressor.WriteCompression = nil
	}

	return compressor, nil
}

// Close is an io.Closer interface function
func (c *Compressor) Close() error {
	if c == nil {
		return nil
	}
	if c.LZMAWriter != nil {
		return c.LZMAWriter.Close()
	}

	return nil
}

// WriteEnabled checks whether write is enabled
func (c *Compressor) WriteEnabled() bool {
	if c == nil {
		return false
	}
	return c.WriteCompression != nil
}

// ReadEnabled checks whether read is enabled
func (c *Compressor) ReadEnabled() bool {
	if c == nil {
		return false
	}
	return c.ReadCompression != nil
}

// Write is an io.Writer function
func (c *Compressor) Write(p []byte) (n int, err error) {
	if c == nil {
		return 0, fmt.Errorf("can't write to empty")
	}

	return c.LZMAWriter.Write(p)
}

// Read is an io.Reader function
func (c *Compressor) Read(p []byte) (n int, err error) {
	if c == nil {
		return 0, fmt.Errorf("can't read from empty")
	}

	return c.LZMAReader.Read(p)
}
