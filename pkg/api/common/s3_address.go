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
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7/pkg/s3utils"
)

// NewS3Address creates new S3Address
func NewS3Address(bucket, object string) *S3Address {
	return &S3Address{
		Bucket: bucket,
		Object: object,
	}
}

// NewS3AddressFromString creates new S3Address from string. The first section of the string is considered to be a bucket
// Expected format is:
// bucket/path/to/the/object
func NewS3AddressFromString(str string) *S3Address {
	parts := strings.SplitN(str, "/", 2)
	if len(parts) != 2 {
		return nil
	}
	bucket := parts[0]
	object := parts[1]
	return NewS3Address(bucket, object)
}

// SetBucket sets bucket
func (x *S3Address) SetBucket(bucket string) *S3Address {
	if x == nil {
		return nil
	}
	x.Bucket = bucket
	return x
}

// SetObject sets object
func (x *S3Address) SetObject(object string) *S3Address {
	if x == nil {
		return nil
	}
	x.Object = object
	return x
}

// GetFilename gets filename of an object
func (x *S3Address) GetFilename() string {
	if x == nil {
		return ""
	}
	return filepath.Base(x.Object)
}

// GetBucketNameError returns an error in case BucketName is incorrect
func (x *S3Address) GetBucketNameError() error {
	if x == nil {
		return nil
	}
	return s3utils.CheckValidBucketNameStrict(x.Bucket)
}

// IsBucketNameValid checks whether bucket name is valid
func (x *S3Address) IsBucketNameValid() bool {
	if x == nil {
		return false
	}
	return x.GetBucketNameError() == nil
}

// GetObjectNameError returns an error in case ObjectName is incorrect
func (x *S3Address) GetObjectNameError() error {
	if x == nil {
		return nil
	}
	return s3utils.CheckValidObjectName(x.Object)
}

// IsObjectNameValid checks whether object name is valid
func (x *S3Address) IsObjectNameValid() bool {
	if x == nil {
		return false
	}
	return x.GetObjectNameError() == nil
}

// IsValid checks whether address is valid
func (x *S3Address) IsValid() bool {
	return x.IsBucketNameValid() && x.IsObjectNameValid()
}

// GetError returns and error in case S3Address is not valid
func (x *S3Address) GetError() error {
	if x == nil {
		return nil
	}

	if x.IsValid() {
		return nil
	}

	return fmt.Errorf("s3Address error. bucket: %v object: %v", x.GetBucketNameError(), x.GetObjectNameError())
}

// String stringifies S3Address
func (x *S3Address) String() string {
	if x != nil {
		return x.Bucket + "/" + x.Object
	}
	return ""
}
