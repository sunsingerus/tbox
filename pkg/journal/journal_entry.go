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

package journal

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sunsingerus/tbox/pkg/api/common"
)

// Entry defines journal entry structure
type Entry struct {
	//
	// Base info tells about the origin of the journal entry
	//

	// Time of the entry/event
	Time time.Time
	// StartTime specifies start time of the action sequence or execution context
	StartTime time.Time

	// EndpointID [MANDATORY] specifies ID of the endpoint (API call handler/Task processor/etc) which produces the entry
	// See EndpointTypeEnum for available options.
	EndpointID int32
	// EndpointInstanceID [OPTIONAL] specifies ID of the particular endpoint instance (ex.: process) which produces the entry
	// Elaborates EndpointID in terms of particular instance specification.
	EndpointInstanceID *common.UUID
	// SourceID [OPTIONAL] specifies ID of the source (possibly external) of the entry
	SourceID *common.UserID
	// ContextID [OPTIONAL] specifies ID of the execution/rpc context associated with the entry
	ContextUID *common.UUID
	// TaskID [OPTIONAL] specifies ID of the task associated with the entry
	TaskUID *common.UUID
	// Type [MANDATORY] specifies type of the entry - what this entry is about.
	// See EntryTypeEnum for available options.
	Type int32

	// Object info tells about object, if any
	// ObjectType specified object type
	// See ObjectTypeEnum for available options.
	ObjectType     int32
	ObjectAddress  *common.Address
	ObjectSize     uint64
	ObjectMetadata *common.Metadata
	ObjectData     []byte

	// Result [OPTIONAL] specifies result of the operation, specified by Type
	Result string
	// Result [OPTIONAL] specifies status of the operation, specified by Type
	Status string
	// Error [OPTIONAL] tells about error encountered during the operation, specified by Type
	Error error
}

// String
func (e *Entry) String() string {
	if e == nil {
		return "this JE is nil"
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Time:%s\n", e.Time)
	_, _ = fmt.Fprintf(b, "StartTime:%s\n", e.StartTime)

	_, _ = fmt.Fprintf(b, "EndpointID:%d\n", e.EndpointID)
	_, _ = fmt.Fprintf(b, "EndpointInstanceID:%d\n", e.EndpointInstanceID)
	_, _ = fmt.Fprintf(b, "SourceID:%s\n", e.SourceID)
	_, _ = fmt.Fprintf(b, "ContextUID:%s\n", e.ContextUID)
	_, _ = fmt.Fprintf(b, "TaskUID:%s\n", e.TaskUID)
	_, _ = fmt.Fprintf(b, "Type:%d\n", e.Type)

	_, _ = fmt.Fprintf(b, "ObjectType:%d\n", e.ObjectType)
	_, _ = fmt.Fprintf(b, "ObjectAddress:%s\n", e.ObjectAddress)
	_, _ = fmt.Fprintf(b, "ObjectSize:%d\n", e.ObjectSize)
	_, _ = fmt.Fprintf(b, "ObjectMetadata:%s\n", e.ObjectMetadata)
	_, _ = fmt.Fprintf(b, "ObjectData:%s\n", e.ObjectData)

	_, _ = fmt.Fprintf(b, "Result:%s\n", e.Result)
	_, _ = fmt.Fprintf(b, "Status:%s\n", e.Status)
	_, _ = fmt.Fprintf(b, "Error:%s\n", e.Error)

	return b.String()
}

// NewEntry
func NewEntry() *Entry {
	return &Entry{}
}

// SetBaseInfo
func (e *Entry) SetBaseInfo(
	start time.Time,
	endpoint int32,
	endpointInstanceID *common.UUID,
	ctxUID *common.UUID,
	taskUID *common.UUID,
	_type int32,
) *Entry {
	if e == nil {
		return nil
	}
	e.Time = time.Now()
	e.StartTime = start
	e.SetEndpointID(endpoint)
	e.SetEndpointInstanceID(endpointInstanceID)
	e.SetCtxUID(ctxUID)
	e.SetTaskUID(taskUID)
	e.SetType(_type)
	return e
}

// SetEndpointID
func (e *Entry) SetEndpointID(endpoint int32) *Entry {
	if e == nil {
		return nil
	}
	e.EndpointID = endpoint
	return e
}

// SetEndpointIInstanceD
func (e *Entry) SetEndpointInstanceID(endpointInstanceID *common.UUID) *Entry {
	if e == nil {
		return nil
	}
	e.EndpointInstanceID = endpointInstanceID
	return e
}

// SetSourceID
func (e *Entry) SetSourceID(userID *common.UserID) *Entry {
	if e == nil {
		return nil
	}
	e.SourceID = userID
	return e
}

// SetCtxID
func (e *Entry) SetCtxUID(ctxUID *common.UUID) *Entry {
	if e == nil {
		return nil
	}
	e.ContextUID = ctxUID
	return e
}

// SetTaskID
func (e *Entry) SetTaskUID(taskUID *common.UUID) *Entry {
	if e == nil {
		return nil
	}
	e.TaskUID = taskUID
	return e
}

// SetType
func (e *Entry) SetType(_type int32) *Entry {
	if e == nil {
		return nil
	}
	e.Type = _type
	return e
}

// SetObject
func (e *Entry) SetObject(
	objectType int32,
	address *common.Address,
	size uint64,
	metadata *common.Metadata,
	data []byte,
) *Entry {
	if e == nil {
		return nil
	}
	e.SetObjectType(objectType)
	e.SetObjectAddress(address)
	e.SetObjectSize(size)
	e.SetObjectMetadata(metadata)
	e.SetObjectData(data)
	return e
}

// SetObjectType
func (e *Entry) SetObjectType(objectType int32) *Entry {
	if e == nil {
		return nil
	}
	e.ObjectType = objectType
	return e
}

// SetObjectAddress
func (e *Entry) SetObjectAddress(address *common.Address) *Entry {
	if e == nil {
		return nil
	}
	e.ObjectAddress = address
	return e
}

// SetObjectSize
func (e *Entry) SetObjectSize(size uint64) *Entry {
	if e == nil {
		return nil
	}
	e.ObjectSize = size
	return e
}

// SetObjectMetadata
func (e *Entry) SetObjectMetadata(metadata *common.Metadata) *Entry {
	if e == nil {
		return nil
	}
	e.ObjectMetadata = metadata
	return e
}

// EnsureObjectMetadata
func (e *Entry) EnsureObjectMetadata() *common.Metadata {
	if e == nil {
		return nil
	}
	if e.ObjectMetadata == nil {
		e.ObjectMetadata = common.NewMetadata()
	}
	return e.ObjectMetadata
}

// SetObjectData
func (e *Entry) SetObjectData(data []byte) *Entry {
	if e == nil {
		return nil
	}
	e.ObjectData = data
	return e
}

// SetResult
func (e *Entry) SetResult(result string) *Entry {
	if e == nil {
		return nil
	}
	e.Result = result
	return e
}

// SetStatus
func (e *Entry) SetStatus(status string) *Entry {
	if e == nil {
		return nil
	}
	e.Status = status
	return e
}

// SetError
func (e *Entry) SetError(err error) *Entry {
	if e == nil {
		return nil
	}
	e.Error = err
	return e
}

// InsertInto inserts entry into a journal
func (e *Entry) InsertInto(j Journaller) {
	j.Insert(e)
}
