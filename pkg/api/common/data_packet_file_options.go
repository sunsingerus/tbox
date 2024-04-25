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

// DataPacketFileOptions describes metadata a.k.a options for data file
type DataPacketFileOptions struct {
	Header   *Metadata
	Metadata *Metadata

	// Compress outgoing data
	Compress bool

	// Decompress incoming data
	Decompress bool
}

// NewDataPacketFileOptions creates new DataChunkFileOptions
func NewDataPacketFileOptions() *DataPacketFileOptions {
	return new(DataPacketFileOptions)
}

// SetHeader is a setter
func (opts *DataPacketFileOptions) SetHeader(header *Metadata) *DataPacketFileOptions {
	if opts == nil {
		return nil
	}
	opts.Header = header
	return opts
}

// GetHeader is a getter
func (opts *DataPacketFileOptions) GetHeader() *Metadata {
	if opts == nil {
		return nil
	}
	return opts.Header
}

// SetMetadata is a setter
func (opts *DataPacketFileOptions) SetMetadata(meta *Metadata) *DataPacketFileOptions {
	if opts == nil {
		return nil
	}
	opts.Metadata = meta
	return opts
}

// GetMetadata is a getter
func (opts *DataPacketFileOptions) GetMetadata() *Metadata {
	if opts == nil {
		return nil
	}
	return opts.Metadata
}

// SetCompress is a setter
func (opts *DataPacketFileOptions) SetCompress(compress bool) *DataPacketFileOptions {
	if opts == nil {
		return nil
	}
	opts.Compress = compress
	return opts
}

// GetCompress is a getter
func (opts *DataPacketFileOptions) GetCompress() bool {
	if opts == nil {
		return false
	}
	return opts.Compress
}

// SetDecompress is a setter
func (opts *DataPacketFileOptions) SetDecompress(decompress bool) *DataPacketFileOptions {
	if opts == nil {
		return nil
	}
	opts.Decompress = decompress
	return opts
}

// GetDecompress is a getter
func (opts *DataPacketFileOptions) GetDecompress() bool {
	if opts == nil {
		return false
	}
	return opts.Decompress
}
