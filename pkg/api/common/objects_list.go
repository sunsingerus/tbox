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

// NewObjectsList
func NewObjectsList() *ObjectsList {
	return new(ObjectsList)
}

// SetStatus
func (x *ObjectsList) SetStatus(status *Status) *ObjectsList {
	if x == nil {
		return nil
	}
	x.Status = status
	return x
}

// AddReport
func (x *ObjectsList) AddReport(reports ...*Report) *ObjectsList {
	if x == nil {
		return nil
	}
	x.Reports = append(x.Reports, reports...)
	return x
}

// AddTask
func (x *ObjectsList) AddTask(tasks ...*Task) *ObjectsList {
	if x == nil {
		return nil
	}
	x.Tasks = append(x.Tasks, tasks...)
	return x
}

// AddStatus
func (x *ObjectsList) AddStatus(statuses ...*Status) *ObjectsList {
	if x == nil {
		return nil
	}
	x.Statuses = append(x.Statuses, statuses...)
	return x
}

// AddObjectStatus
func (x *ObjectsList) AddObjectStatus(statuses ...*ObjectStatus) *ObjectsList {
	if x == nil {
		return nil
	}
	x.ObjectStatuses = append(x.ObjectStatuses, statuses...)
	return x
}

// AddFile
func (x *ObjectsList) AddFile(files ...*File) *ObjectsList {
	if x == nil {
		return nil
	}
	x.Files = append(x.Files, files...)
	return x
}

// LenReports
func (x *ObjectsList) LenReports() int {
	if x == nil {
		return 0
	}
	return len(x.Reports)
}

// LenTasks
func (x *ObjectsList) LenTasks() int {
	if x == nil {
		return 0
	}
	return len(x.Tasks)
}

// LenStatuses
func (x *ObjectsList) LenStatuses() int {
	if x == nil {
		return 0
	}
	return len(x.Statuses)
}

// LenObjectStatuses
func (x *ObjectsList) LenObjectStatuses() int {
	if x == nil {
		return 0
	}
	return len(x.ObjectStatuses)
}

// LenFiles
func (x *ObjectsList) LenFiles() int {
	if x == nil {
		return 0
	}
	return len(x.Files)
}

/*
// First
func (x *ReportMulti) First() *Report {
	if x.Len() > 0 {
		return x.Reports[0]
	}
	return nil
}

// Shift
func (x *ReportMulti) Shift() *Report {
	if x.Len() > 0 {
		r := x.Reports[0]
		x.Reports = x.Reports[1:]
		return r
	}
	return nil
}
*/

// String
func (x *ObjectsList) String() string {
	return "to be implemented"
}
