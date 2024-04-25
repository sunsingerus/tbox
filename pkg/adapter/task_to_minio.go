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

package adapter

import (
	"fmt"
	"path/filepath"

	minioGo "github.com/minio/minio-go/v7"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/minio"
)

// TaskMinIO specifies task to MinIO adapter
type TaskMinIO struct {
	Config sections.MinIOConfigurator
	Task   *common.Task
}

// Interface compliance check
var (
	_ ObjectToMinIOAdapter = &TaskMinIO{}
)

// NewTaskMinIOAdapter creates new Task to MinIO adapter
func NewTaskMinIOAdapter(config sections.MinIOConfigurator, task *common.Task) *TaskMinIO {
	return &TaskMinIO{
		Config: config,
		Task:   task,
	}
}

// Folder specifies folders of the adapter.
// Base functions operate with any folder, thus custom folders can be used as well
type Folder string

// List of pre-defined folders, typically used in most projects
const (
	task   Folder = "task"
	in     Folder = "in"
	out    Folder = "out"
	result Folder = "result"
	tmp    Folder = "tmp"
)

// GetBucket is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) GetBucket() string {
	return a.Config.GetMinIOBucket()
}

// GetPath is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) GetPath(path ...Folder) string {
	if len(path) > 0 {
		switch path[0] {
		case task:
			return a.Task.GetUuid().String()
		case in:
			return minio.PathJoin(a.GetTaskPath(), "in")
		case out:
			return minio.PathJoin(a.GetTaskPath(), "out")
		case result:
			return minio.PathJoin(a.GetTaskPath(), "result")
		case tmp:
			return minio.PathJoin(a.GetTaskPath(), "tmp")
		}
	}
	return ""
}

// GetPathAddress is a wrapper over GetPath to get Path as an S3Address
func (a *TaskMinIO) GetPathAddress(path ...Folder) *common.S3Address {
	return common.NewS3Address(a.GetBucket(), a.GetPath(path...))
}

// GetPrefix is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) GetPrefix(path ...Folder) string {
	if len(path) > 0 {
		switch path[0] {
		case task:
			return a.GetTaskPath() + "/"
		case in:
			return a.GetInPath() + "/"
		case out:
			return a.GetOutPath() + "/"
		case result:
			return a.GetResultPath() + "/"
		case tmp:
			return a.GetTmpPath() + "/"
		}
	}
	return ""
}

// GetObjectPath is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) GetObjectPath(object string, base bool, path ...Folder) string {
	if len(path) > 0 {
		if base {
			object = filepath.Base(object)
		}
		switch path[0] {
		case task:
			return minio.PathJoin(a.GetTaskPath(), object)
		case in:
			return minio.PathJoin(a.GetInPath(), object)
		case out:
			return minio.PathJoin(a.GetOutPath(), object)
		case result:
			return minio.PathJoin(a.GetResultPath(), object)
		case tmp:
			return minio.PathJoin(a.GetTmpPath(), object)
		}
	}
	return ""
}

// GetObjectAddress is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) GetObjectAddress(object string, base bool, path ...Folder) *common.S3Address {
	return common.NewS3Address(a.GetBucket(), a.GetObjectPath(object, base, path...))
}

// FirstObject is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) FirstObject(path ...Folder) (*common.S3Address, error) {
	mi, err := minio.NewMinIOFromConfig(a.Config)
	if err != nil {
		return nil, err
	}
	bucket := a.GetBucket()
	list, err := mi.List(bucket, a.GetPrefix(path...), 1)
	if err != nil {
		return nil, err
	}

	if len(list) != 1 {
		return nil, fmt.Errorf("not found")
	}

	info := list[0]
	return common.NewS3Address(bucket, info.Key), nil
}

// WalkObjects is an ObjectToMinIOAdapter interface function
func (a *TaskMinIO) WalkObjects(f func(int, *minio.MinIO, *minioGo.ObjectInfo, *common.S3Address) bool, path ...Folder) error {
	mi, err := minio.NewMinIOFromConfig(a.Config)
	if err != nil {
		return err
	}
	bucket := a.GetBucket()
	list, err := mi.List(bucket, a.GetPrefix(path...), -1)
	if err != nil {
		return err
	}

	for i := range list {
		info := list[i]
		s3address := common.NewS3Address(bucket, info.Key)
		if !f(len(list), mi, &info, s3address) {
			break
		}
	}
	return nil
}

/***********************/
/*     Wrappers        */
/***********************/

// GetTaskPath gets path of the whole task
func (a *TaskMinIO) GetTaskPath() string {
	return a.GetPath(task)
}

// GetInPath gets `in` path
func (a *TaskMinIO) GetInPath() string {
	return a.GetPath(in)
}

// GetInPathAddress gets `in` path as S3Address
func (a *TaskMinIO) GetInPathAddress() *common.S3Address {
	return a.GetPathAddress(in)
}

// GetInPrefix gets `in` prefix
func (a *TaskMinIO) GetInPrefix() string {
	return a.GetPrefix(in)
}

// GetInFile gets `in` full file path
func (a *TaskMinIO) GetInFile(file string, base bool) string {
	return a.GetObjectPath(file, base, in)
}

// GetInFileAddress gets `in` file S3Address
func (a *TaskMinIO) GetInFileAddress(file string, base bool) *common.S3Address {
	return a.GetObjectAddress(file, base, in)
}

// FirstInFile gets first `in` file S3Address
func (a *TaskMinIO) FirstInFile() (*common.S3Address, error) {
	return a.FirstObject(in)
}

// WalkInFiles walks `in` files
func (a *TaskMinIO) WalkInFiles(f func(int, *minio.MinIO, *minioGo.ObjectInfo, *common.S3Address) bool) error {
	return a.WalkObjects(f, in)
}

// GetOutPath gets `out` path
func (a *TaskMinIO) GetOutPath() string {
	return a.GetPath(out)
}

// GetOutPathAddress gets `out` path as S3Address
func (a *TaskMinIO) GetOutPathAddress() *common.S3Address {
	return a.GetPathAddress(out)
}

// GetOutPrefix gets `out` prefix
func (a *TaskMinIO) GetOutPrefix() string {
	return a.GetPrefix(out)
}

// GetOutFile gets `out` full file path
func (a *TaskMinIO) GetOutFile(file string, base bool) string {
	return a.GetObjectPath(file, base, out)
}

// GetOutFileAddress gets `out` file S3Address
func (a *TaskMinIO) GetOutFileAddress(file string, base bool) *common.S3Address {
	return a.GetObjectAddress(file, base, out)
}

// FirstOutFile gets first `out` file S3Address
func (a *TaskMinIO) FirstOutFile() (*common.S3Address, error) {
	return a.FirstObject(out)
}

// WalkOutFiles walks `out` files
func (a *TaskMinIO) WalkOutFiles(f func(int, *minio.MinIO, *minioGo.ObjectInfo, *common.S3Address) bool) error {
	return a.WalkObjects(f, out)
}

// GetResultPath gets `result` path
func (a *TaskMinIO) GetResultPath() string {
	return a.GetPath(result)
}

// GetResultPathAddress gets `result` path as S3Address
func (a *TaskMinIO) GetResultPathAddress() *common.S3Address {
	return a.GetPathAddress(result)
}

// GetResultPrefix gets `result` prefix
func (a *TaskMinIO) GetResultPrefix() string {
	return a.GetPrefix(result)
}

// GetResultFile gets `result` file path
func (a *TaskMinIO) GetResultFile(file string, base bool) string {
	return a.GetObjectPath(file, base, result)
}

// GetResultFileAddress gets `result` file S3Address
func (a *TaskMinIO) GetResultFileAddress(file string, base bool) *common.S3Address {
	return a.GetObjectAddress(file, base, result)
}

// FirstResultFile gets first `result` file S3Address
func (a *TaskMinIO) FirstResultFile() (*common.S3Address, error) {
	return a.FirstObject(result)
}

// WalkResultFiles walks `result` files
func (a *TaskMinIO) WalkResultFiles(f func(int, *minio.MinIO, *minioGo.ObjectInfo, *common.S3Address) bool) error {
	return a.WalkObjects(f, result)
}

// GetTmpPath gets `tmp` path
func (a *TaskMinIO) GetTmpPath() string {
	return a.GetPath(tmp)
}

// GetTmpPathAddress gets `tmp` path as S3Address
func (a *TaskMinIO) GetTmpPathAddress() *common.S3Address {
	return a.GetPathAddress(tmp)
}

// GetTmpPrefix gets `tmp` prefix
func (a *TaskMinIO) GetTmpPrefix() string {
	return a.GetPrefix(tmp)
}

// GetTmpFile gets `tmp` file path
func (a *TaskMinIO) GetTmpFile(file string, base bool) string {
	return a.GetObjectPath(file, base, tmp)
}

// GetTmpFileAddress gets `tmp` files S3Address
func (a *TaskMinIO) GetTmpFileAddress(file string, base bool) *common.S3Address {
	return a.GetObjectAddress(file, base, tmp)
}

// FirstTmpFile gets first `tmp` file S3Address
func (a *TaskMinIO) FirstTmpFile() (*common.S3Address, error) {
	return a.FirstObject(tmp)
}

// WalkTmpFiles walks `tmp` files
func (a *TaskMinIO) WalkTmpFiles(f func(int, *minio.MinIO, *minioGo.ObjectInfo, *common.S3Address) bool) error {
	return a.WalkObjects(f, tmp)
}
