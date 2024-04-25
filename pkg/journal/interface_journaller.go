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
	"github.com/sunsingerus/tbox/pkg/api/common"
)

// Journaller
type Journaller interface {
	SetContext(ctx Contexter) Journaller
	SetTask(task Tasker) Journaller
	WithContext(ctx Contexter) Journaller
	WithTask(task Tasker) Journaller

	//
	// Expose direct access to storage via adapters.
	// Implement Adapter interface as wrappers over Adapter
	//

	NewEntry(action int32) *Entry
	Adapter

	//
	// Common requests section
	//

	RequestStart()
	RequestEnd()
	RequestError(callErr error)

	//
	// In-request actions
	//

	SaveData(
		address *common.Address,
		size int64,
		metadata *common.Metadata,
		data []byte,
	)
	SaveDataError(callErr error)

	//
	//
	//

	ProcessData(
		address *common.Address,
		size int64,
		metadata *common.Metadata,
	)
	ProcessDataError(callErr error)

	//
	//
	//

	Result(
		address *common.Address,
		size int64,
		metadata *common.Metadata,
	)

	//
	//
	//

	SaveTask(task *common.Task)
	SaveTaskError(task *common.Task, callErr error)

	//
	//
	//

	ProcessTask(task *common.Task)
	ProcessTaskError(task *common.Task, callErr error)

	//
	//
	//
	Lookup(address *common.Address)
	LookupError(address *common.Address, callErr error)
}
