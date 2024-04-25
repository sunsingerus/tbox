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

const (
	// DataChunkTypeReserved is a reserved data chunk type. Due to first enum value has to be zero in proto3.
	DataChunkTypeReserved int32 = 0
	// DataChunkTypeUnspecified is an unspecified data chunk type. Unspecified means data chunk type is unknown.
	DataChunkTypeUnspecified int32 = 100
	// DataChunkTypeData is a custom data Data chunk.
	DataChunkTypeData int32 = 200
)

// DataChunkTypeEnum is an enum of available data chunk types.
var DataChunkTypeEnum = NewEnum()

func init() {
	DataChunkTypeEnum.MustRegister("DataChunkTypeReserved", DataChunkTypeReserved)
	DataChunkTypeEnum.MustRegister("DataChunkTypeUnspecified", DataChunkTypeUnspecified)
	DataChunkTypeEnum.MustRegister("DataChunkTypeData", DataChunkTypeData)
}

// NewDataPacket creates new DataPacket.
func NewDataPacket() *DataPacket {
	return new(DataPacket)
}

//
// Interface functions
//

// GetDataLen is an IDataChunk interface function
func (x *DataPacket) GetDataLen() int {
	return x.GetDataChunk().GetDataLen()
}

// GetData is an IDataChunk interface function
func (x *DataPacket) GetData() []byte {
	return x.GetDataChunk().GetData()
}

// SetData is an IDataChunk interface function
func (x *DataPacket) SetData(data []byte) {
	x.EnsureDataChunk().SetData(data)
}

// GetLast is an IDataChunk interface function
func (x *DataPacket) GetLast() bool {
	return x.GetDataChunk().GetLast()
}

// SetLast is an IDataChunk interface function
func (x *DataPacket) SetLast(last bool) {
	x.EnsureDataChunk().SetLast(last)
}

// GetOffset is an IDataChunk interface function
func (x *DataPacket) GetOffset() int64 {
	return x.GetDataChunk().GetOffset()
}

// SetOffset is an IDataChunk interface function
func (x *DataPacket) SetOffset(offset int64) {
	x.EnsureDataChunk().SetOffset(offset)
}

// GetIDataChunk is an IDataChunkEnvelope interface function
func (x *DataPacket) GetIDataChunk() IDataChunk {
	return x.GetDataChunk()
}

// EnsureIDataChunk is an IDataChunkEnvelope interface function
func (x *DataPacket) EnsureIDataChunk() IDataChunk {
	return x.EnsureDataChunk()
}

// SetDataChunk is a setter
func (x *DataPacket) SetDataChunk(chunk *DataChunk) {
	if x == nil {
		return
	}
	x.DataChunk = chunk
}

// EnsureDataChunk is a setter with guaranteed result
func (x *DataPacket) EnsureDataChunk() *DataChunk {
	if x.GetDataChunk() == nil {
		x.SetDataChunk(NewDataChunk())
	}
	return x.GetDataChunk()
}

// SetStreamOptions is a setter
func (x *DataPacket) SetStreamOptions(o *PresentationOptions) {
	if x == nil {
		return
	}
	x.StreamOptions = o
}

// EnsureStreamOptions is a getter with guaranteed result
func (x *DataPacket) EnsureStreamOptions() *PresentationOptions {
	if x.GetStreamOptions() == nil {
		x.SetStreamOptions(NewPresentationOptions())
	}
	return x.GetStreamOptions()
}

// SetPayloadMetadata is a setter
func (x *DataPacket) SetPayloadMetadata(m *Metadata) {
	if x == nil {
		return
	}
	x.PayloadMetadata = m
}

// String make string representation.
func (x *DataPacket) String() string {
	return "no be implemented"
}
