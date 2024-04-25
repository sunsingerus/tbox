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
	minio_go "github.com/minio/minio-go/v7"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/minio"
)

// ObjectToMinIOAdapter is an interface of an entity which can adapt an object to MinIO storage
type ObjectToMinIOAdapter interface {
	// GetBucket returns bucket within the storage
	GetBucket() string
	// GetPath returns specified task-related path within the storage.
	// Path does not end with / and thus can't be used as prefix for uploading objects.
	GetPath(path ...Folder) string
	// GetPrefix returns specified task-related prefix within storage.
	// Prefix ends with / and thus can be used as prefix for uploading objects.
	GetPrefix(path ...Folder) string
	// GetObjectPath returns specified task-related full path to an object (file) within the storage.
	// Filename is included into the result. Can be used to upload/download an object by its full path.
	// In case `base` specified, object is trimmed to base - maning filename used w/o prefix.
	GetObjectPath(object string, base bool, path ...Folder) string
	// GetObjectAddress the same idea as GetObjectPath, but returns ready-to-use S3Address
	GetObjectAddress(object string, base bool, path ...Folder) *common.S3Address
	// FirstObject returns specified task-related full path to the first object (file) (filename included) within storage.
	// Can be used to upload/download the object by full path.
	// Is the same as GetObjectPath with the difference that it makes GetObjectPath() of the first object in the path
	FirstObject(path ...Folder) (*common.S3Address, error)
	// WalkObjects walks files with function.
	WalkObjects(f func(int, *minio.MinIO, *minio_go.ObjectInfo, *common.S3Address) bool, path ...Folder) error
}
