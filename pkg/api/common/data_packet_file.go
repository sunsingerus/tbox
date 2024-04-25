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
	log "github.com/sirupsen/logrus"
)

// DataPacketFile is a handler to set of DataChunk(s)
type DataPacketFile struct {
	*DataChunkFile

	// writer provides functions to send DataPacket (envelope) with DataChunk inside
	writer DataPacketWriter
	// reader provides functions to receive DataPacket (envelope) with DataChunk inside
	reader DataPacketReader

	// payloadMetadata is optional.
	payloadMetadata *Metadata
	// streamOptions is optional.
	streamOptions *PresentationOptions
}

// Ensure interface compatibility
var (
	_ IDataChunkFile      = &DataPacketFile{}
	_ IDataChunkTransport = &DataPacketFile{}
)

// OpenDataPacketFile opens set of DataChunk(s)
// Inspired by os.OpenFile()
func OpenDataPacketFile(writer DataPacketWriter, reader DataPacketReader) (f *DataPacketFile, err error) {
	f = &DataPacketFile{
		writer: writer,
		reader: reader,
	}
	if f.DataChunkFile, err = OpenDataChunkFile(f); err != nil {
		return nil, err
	}

	return f, nil
}

// Recv is an IDataChunkTransport interface function
func (f *DataPacketFile) Recv() (IDataChunk, error) {
	packet, err := f.reader.Recv()
	f.acceptPayloadMetadata(packet)
	f.acceptStreamOptions(packet)
	return packet, err
}

// Send is an IDataChunkTransport interface function
func (f *DataPacketFile) Send(iDataChunk IDataChunk) error {
	// Send DataPacket (envelope) to the writer
	packet := iDataChunk.(*DataPacket)
	return f.writer.Send(packet)
}

/*
func S(){
	// Fetch filename from the chunks stream - it may be in any chunk, actually
	filename := f.PayloadMetadata.GetFilename()
	if filename == "" {
		filename = "not specified"
	}
}
*/

// NewIDataChunk is an IDataChunkTransport interface function and create abstracted IDataChunk
func (f *DataPacketFile) NewIDataChunk(offsetter GetOffsetter) IDataChunk {
	packet := NewDataPacket()
	if offsetter.GetOffset() == 0 {
		// First chunk in this file, it may have some aux data
		if f.GetPayloadMetadata() != nil {
			packet.SetPayloadMetadata(f.GetPayloadMetadata())
		}
		if f.GetStreamOptions() != nil {
			packet.SetStreamOptions(f.GetStreamOptions())
		}
		log.Tracef("Attaching metadata. payload=%v", f.payloadMetadata)
	}
	return packet
}

// acceptPayloadMetadata accepts payload metadata from DataChunk into the file
func (f *DataPacketFile) acceptPayloadMetadata(packet *DataPacket) {
	if packet.GetPayloadMetadata() != nil {
		f.SetPayloadMetadata(packet.GetPayloadMetadata())
	}
}

// acceptStreamOptions accepts stream options from DataChunk into the file
func (f *DataPacketFile) acceptStreamOptions(packet *DataPacket) {
	if packet.GetStreamOptions() != nil {
		f.SetStreamOptions(packet.GetStreamOptions())
	}
}

// GetPayloadMetadata is an aux function
func (f *DataPacketFile) GetPayloadMetadata() *Metadata {
	if f == nil {
		return nil
	}
	return f.payloadMetadata
}

// EnsurePayloadMetadata is an aux function
func (f *DataPacketFile) EnsurePayloadMetadata() *Metadata {
	if f.GetPayloadMetadata() == nil {
		f.SetPayloadMetadata(NewMetadata())
	}
	return f.GetPayloadMetadata()
}

// SetPayloadMetadata is an aux function
func (f *DataPacketFile) SetPayloadMetadata(payload *Metadata) {
	if f == nil {
		return
	}
	f.payloadMetadata = payload
}

// GetStreamOptions is an aux function
func (f *DataPacketFile) GetStreamOptions() *PresentationOptions {
	if f == nil {
		return nil
	}
	return f.streamOptions
}

// EnsureStreamOptions is an aux function
func (f *DataPacketFile) EnsureStreamOptions() *PresentationOptions {
	if f.GetStreamOptions() == nil {
		f.SetStreamOptions(NewPresentationOptions())
	}
	return f.GetStreamOptions()
}

// SetStreamOptions is an aux function
func (f *DataPacketFile) SetStreamOptions(opts *PresentationOptions) {
	if f == nil {
		return
	}
	f.streamOptions = opts
}

// GetFilename returns filename.
// Shortcut to get filename from metadata
func (f *DataPacketFile) GetFilename() string {
	return f.GetPayloadMetadata().GetFilename()
}

// Name is a wrapper for GetFilename()
// Main purpose is to have the same Name() function as os.File does (to be used as interface function)
func (f *DataPacketFile) Name() string {
	return f.GetFilename()
}

// SetFilename sets filename.
// Shortcut to set filename from metadata
func (f *DataPacketFile) SetFilename(filename string) *DataPacketFile {
	f.EnsurePayloadMetadata().SetFilename(filename)
	return f
}
