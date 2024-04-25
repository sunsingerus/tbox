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
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	// Due to first enum value has to be zero in proto3
	TaskReserved int32 = 0
	// Unspecified
	TaskUnspecified int32 = 100
	// Echo request expects echo reply as an answer
	TaskEchoRequest int32 = 200
	// Echo reply is an answer to echo request
	TaskEchoReply int32 = 300
	// Request for configuration from the other party
	TaskConfigRequest int32 = 400
	// Configuration
	TaskConfig int32 = 500
	// Metrics schedule sends schedule by which metrics should be sent.
	TaskMetricsSchedule int32 = 600
	// Metrics request is an explicit request for metrics to be sent
	TaskMetricsRequest int32 = 700
	// One-time metrics
	TaskMetrics int32 = 800
	// Schedule to send data
	TaskDataSchedule int32 = 900
	// Explicit data request
	TaskDataRequest int32 = 1000
	// Data are coming
	TaskData int32 = 1100
	// Address is coming
	TaskAddress            int32 = 1200
	TaskExtract            int32 = 1300
	TaskExtractExecutables int32 = 1400
)

var TaskTypeEnum = NewEnum()

func init() {
	TaskTypeEnum.MustRegister("TaskReserved", TaskReserved)
	TaskTypeEnum.MustRegister("TaskUnspecified", TaskUnspecified)
	TaskTypeEnum.MustRegister("TaskEchoRequest", TaskEchoRequest)
	TaskTypeEnum.MustRegister("TaskEchoReply", TaskEchoReply)
	TaskTypeEnum.MustRegister("TaskConfigRequest", TaskConfigRequest)
	TaskTypeEnum.MustRegister("TaskConfig", TaskConfig)
	TaskTypeEnum.MustRegister("TaskMetricsSchedule", TaskMetricsSchedule)
	TaskTypeEnum.MustRegister("TaskMetricsRequest", TaskMetricsRequest)
	TaskTypeEnum.MustRegister("TaskMetrics", TaskMetrics)
	TaskTypeEnum.MustRegister("TaskDataSchedule", TaskDataSchedule)
	TaskTypeEnum.MustRegister("TaskDataRequest", TaskDataRequest)
	TaskTypeEnum.MustRegister("TaskData", TaskData)
	TaskTypeEnum.MustRegister("TaskAddress", TaskAddress)
	TaskTypeEnum.MustRegister("TaskExtract", TaskExtract)
	TaskTypeEnum.MustRegister("TaskExtractExecutables", TaskExtractExecutables)
}

// NewTask creates new Command with pre-allocated header
func NewTask() *Task {
	return &Task{
		Header: NewMetadata(),
	}
}

// NewTaskUnmarshalFrom creates new Task from a slice of bytes
func NewTaskUnmarshalFrom(bytes []byte) (*Task, error) {
	task := new(Task)
	if err := task.UnmarshalFrom(bytes); err != nil {
		return nil, err
	}
	return task, nil
}

// UnmarshalFrom unmarshal commands from a slice of bytes
func (x *Task) UnmarshalFrom(bytes []byte) error {
	return proto.Unmarshal(bytes, x)
}

// EnsureHeader
func (x *Task) EnsureHeader() *Metadata {
	if x == nil {
		return nil
	}
	if x.Header == nil {
		x.Header = NewMetadata()
	}
	return x.Header
}

// SetBytes sets bytes as task's data. Provided bytes are not interpreted and used as-is.
func (x *Task) SetBytes(bytes []byte) *Task {
	x.Bytes = bytes
	return x
}

// SetPayload puts any protobuf message (type) into task's data.
// Message is marshalled into binary form and set as data bytes of the task.
func (x *Task) SetPayload(msg proto.Message) error {
	if bytes, err := proto.Marshal(msg); err == nil {
		x.SetBytes(bytes)
		return nil
	} else {
		return err
	}
}

// GetPayload extracts profobuf message (type) from task's data.
// Message is unmarshalled from task's data into provided message.
func (x *Task) GetPayload(msg proto.Message) error {
	return proto.Unmarshal(x.GetBytes(), msg)
}

// AddSubject ands one subject to the task
func (x *Task) AddSubject(subject *Metadata) *Task {
	x.Subjects = append(x.Subjects, subject)
	return x
}

// AddSubjects adds multiple subjects to the task
func (x *Task) AddSubjects(subjects ...*Metadata) *Task {
	x.Subjects = append(x.Subjects, subjects...)
	return x
}

// FirstSubject fetches first (0-indexed) subject. List of list does not change.
func (x *Task) FirstSubject() *Metadata {
	if x == nil {
		return nil
	}
	if len(x.Subjects) == 0 {
		return nil
	}
	return x.Subjects[0]
}

// LastSubject fetches last subject. List of list does not change.
func (x *Task) LastSubject() *Metadata {
	if x == nil {
		return nil
	}
	if len(x.Subjects) == 0 {
		return nil
	}
	return x.Subjects[len(x.Subjects)-1]
}

// AddSubtask adds one subtask to the task
func (x *Task) AddSubtask(task *Task) *Task {
	x.Children = append(x.Children, task)
	return x
}

// AddSubtasks adds multiple subtasks to the task
func (x *Task) AddSubtasks(tasks ...*Task) *Task {
	x.Children = append(x.Children, tasks...)
	return x
}

// FirstSubtask fetches first (0-indexed) subtask. List of subtasks does not change.
func (x *Task) FirstSubtask() *Task {
	if x == nil {
		return nil
	}
	if len(x.Children) == 0 {
		return nil
	}
	return x.Children[0]
}

// LastSubtask fetches last subtask. List of subtasks does not change.
func (x *Task) LastSubtask() *Task {
	if x == nil {
		return nil
	}
	if len(x.Children) == 0 {
		return nil
	}
	return x.Children[len(x.Children)-1]
}

// ShiftSubtasks fetches first (0-indexed) task from available tasks.
// Fetched task is removed from the list of tasks.
// List of subtasks changes.
func (x *Task) ShiftSubtasks() *Task {
	var task *Task = nil
	if len(x.Children) > 0 {
		task = x.Children[0]
		x.Children = x.Children[1:]
	}
	return task
}

// Derive produces derivative task from the task as:
//  1. fetches first (0-indexed) subtask from available subtasks
//  2. and attaches all the rest subtasks (if any) as subtasks of the fetched one, which it the new top now.
//
// Original task is modified.
func (x *Task) Derive() *Task {
	// Assume new root task is the first subtask of current task
	root := x.FirstSubtask()
	if root == nil {
		return nil
	}

	// Parent of the new task is current task
	root.AddParent(x)
	// Subtasks of the new task are the same as of the current task except the rrot itself
	root.Children = x.Children
	root.ShiftSubtasks()
	return root
}

// AddParent adds one parent of the task
func (x *Task) AddParent(task *Task) *Task {
	x.Parents = append(x.Parents, task)
	return x
}

// AddParents adds multiple parents of the task
func (x *Task) AddParents(tasks ...*Task) *Task {
	x.Parents = append(x.Parents, tasks...)
	return x
}

// FirstParent fetches first (0-indexed) parent. List of list does not change.
func (x *Task) FirstParent() *Task {
	if x == nil {
		return nil
	}
	if len(x.Parents) == 0 {
		return nil
	}
	return x.Parents[0]
}

// LastParent fetches last parent. List of list does not change.
func (x *Task) LastParent() *Task {
	if x == nil {
		return nil
	}
	if len(x.Parents) == 0 {
		return nil
	}
	return x.Parents[len(x.Parents)-1]
}

// String
func (x *Task) String() string {
	if x == nil {
		return "nil"
	}

	var parts []string
	if _type := x.GetType(); _type > 0 {
		parts = append(parts, "type:"+fmt.Sprintf("%d", _type))
	}
	if name := x.GetName(); name != "" {
		parts = append(parts, "name:"+name)
	}
	if len(x.GetSubjects()) > 0 {
		parts = append(parts, "subjects:")
		for _, subj := range x.GetSubjects() {
			parts = append(parts, subj.String())
		}
	}

	if len(x.Children) > 0 {
		parts = append(parts, fmt.Sprintf("%d children", len(x.Children)))
	}

	if len(x.Parents) > 0 {
		parts = append(parts, fmt.Sprintf("%d parents", len(x.Parents)))
	}

	return strings.Join(parts, " ")
}
