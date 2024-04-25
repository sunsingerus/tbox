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
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
)

// BaseJournal
type BaseJournal struct {
	start              time.Time
	endpointID         int32
	endpointInstanceID *common.UUID
	adapter            Adapter

	NOPJournal
}

// Validate interface compatibility
var _ Journaller = &BaseJournal{}

// NewBaseJournal
func NewBaseJournal(endpointID int32, endpointInstanceID *common.UUID, adapter Adapter) (*BaseJournal, error) {
	return &BaseJournal{
		start:              time.Now(),
		endpointID:         endpointID,
		endpointInstanceID: endpointInstanceID,
		adapter:            adapter,
	}, nil
}

// copy
func (j *BaseJournal) copy() *BaseJournal {
	return &BaseJournal{
		start:              j.start,
		endpointID:         j.endpointID,
		endpointInstanceID: j.endpointInstanceID,
		adapter:            j.adapter,
	}
}

// SetContext
func (j *BaseJournal) SetContext(ctx Contexter) Journaller {
	fmt.Println(fmt.Sprintf("SetContext. UUID=%s\n", ctx.GetUuid()))
	if j == nil {
		return nil
	}
	j.ctx = ctx
	return j
}

// GetContext
func (j *BaseJournal) GetContext() Contexter {
	if j == nil {
		return nil
	}
	return j.ctx
}

// GetContextUUID
func (j *BaseJournal) GetContextUUID() *common.UUID {
	if j.GetContext() == nil {
		return nil
	}
	return j.GetContext().GetUuid()
}

// SetTask
func (j *BaseJournal) SetTask(task Tasker) Journaller {
	fmt.Println(fmt.Sprintf("SetTask. UUID=%s\n", task.GetUuid()))
	if j == nil {
		return nil
	}
	j.task = task
	return j
}

// GetTask
func (j *BaseJournal) GetTask() Tasker {
	if j == nil {
		return nil
	}
	return j.task
}

// GetTaskUUID
func (j *BaseJournal) GetTaskUUID() *common.UUID {
	if j.GetTask() == nil {
		return nil
	}
	return j.GetTask().GetUuid()
}

// WithContext
func (j *BaseJournal) WithContext(ctx Contexter) Journaller {
	return j.copy().SetContext(ctx)
}

// WithTask
func (j *BaseJournal) WithTask(task Tasker) Journaller {
	return j.copy().SetTask(task)
}

// NewEntry
func (j *BaseJournal) NewEntry(entryType int32) *Entry {
	return NewEntry().SetBaseInfo(j.start, j.endpointID, j.endpointInstanceID, j.GetContextUUID(), j.GetTaskUUID(), entryType)
}

// Insert
func (j *BaseJournal) Insert(entry *Entry) error {
	if j == nil {
		return fmt.Errorf("unable to unsert into nil")
	}
	err := j.adapter.Insert(entry)
	if err != nil {
		log.Warnf("unable to insert journal entry")
	}
	return err
}

// FindAll
func (j *BaseJournal) FindAll(entry *Entry) ([]*Entry, error) {
	if j == nil {
		return nil, fmt.Errorf("unable to find over nil")
	}
	return j.adapter.FindAll(entry)
}

// RequestStart journals beginning of the request processing
func (j *BaseJournal) RequestStart() {
	e := j.NewEntry(EntryTypeRequestStart)
	_ = j.Insert(e)
}

// RequestEnd journals request completed successfully
func (j *BaseJournal) RequestEnd() {
	e := j.NewEntry(EntryTypeRequestCompleted)
	_ = j.Insert(e)
}

// RequestError journals request has failed with an error
func (j *BaseJournal) RequestError(callErr error) {
	e := j.NewEntry(EntryTypeRequestError).SetError(callErr)
	_ = j.Insert(e)
}

// SaveData journals data saved successfully
func (j *BaseJournal) SaveData(
	address *common.Address,
	size int64,
	metadata *common.Metadata,
	data []byte,
) {
	e := j.NewEntry(EntryTypeSaveData).
		SetSourceID(metadata.GetUserID()).
		SetObject(metadata.GetType(), address, uint64(size), metadata, data)
	_ = j.Insert(e)
}

// SaveDataError journals data not saved due to an error
func (j *BaseJournal) SaveDataError(callErr error) {
	e := j.NewEntry(EntryTypeSaveDataError).SetError(callErr)
	_ = j.Insert(e)
}

// ProcessData journals data processed successfully
func (j *BaseJournal) ProcessData(
	address *common.Address,
	size int64,
	metadata *common.Metadata,
) {
	e := j.NewEntry(EntryTypeProcessData).
		SetSourceID(metadata.GetUserID()).
		SetObject(metadata.GetType(), address, uint64(size), metadata, nil)
	_ = j.Insert(e)
}

// ProcessDataError journals data not processed due to an error
func (j *BaseJournal) ProcessDataError(callErr error) {
	e := j.NewEntry(EntryTypeProcessDataError).SetError(callErr)
	_ = j.Insert(e)
}

// Result journals result
func (j *BaseJournal) Result(
	address *common.Address,
	size int64,
	metadata *common.Metadata,
) {
	e := j.NewEntry(EntryTypeResult).
		SetSourceID(metadata.GetUserID()).
		SetObject(metadata.GetType(), address, uint64(size), metadata, nil)
	_ = j.Insert(e)
}

// SaveTask journals task saved successfully
func (j *BaseJournal) SaveTask(task *common.Task) {
	e := j.NewEntry(EntryTypeSaveTask).SetTaskUID(task.GetUuid())
	_ = j.Insert(e)
}

// SaveTaskError journals task not saved due to an error
func (j *BaseJournal) SaveTaskError(task *common.Task, callErr error) {
	e := j.NewEntry(EntryTypeSaveTaskError).SetError(callErr).SetTaskUID(task.GetUuid())
	_ = j.Insert(e)
}

// ProcessTask journals task processed successfully
func (j *BaseJournal) ProcessTask(task *common.Task) {
	//e := j.NewEntry(EntryTypeProcessTask).SetTaskID(task.GetUUID()).SetObjectAddress(task.GetResult())
	e := j.NewEntry(EntryTypeProcessTask).SetTaskUID(task.GetUuid())
	_ = j.Insert(e)
}

// ProcessTaskError journals task not processed due to an error
func (j *BaseJournal) ProcessTaskError(task *common.Task, callErr error) {
	e := j.NewEntry(EntryTypeProcessTaskError).SetError(callErr).SetTaskUID(task.GetUuid())
	_ = j.Insert(e)
}

// Lookup
func (j *BaseJournal) Lookup(address *common.Address) {
	e := j.NewEntry(EntryTypeLookup).SetObjectAddress(address)
	_ = j.Insert(e)
}

// LookupError
func (j *BaseJournal) LookupError(address *common.Address, callErr error) {
	e := j.NewEntry(EntryTypeLookupError).SetError(callErr).SetObjectAddress(address)
	_ = j.Insert(e)
}
