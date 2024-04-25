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

package sections

import (
	"fmt"

	"github.com/sunsingerus/tbox/pkg/config/items"
)

// MinIOConfigurator
type MinIOConfigurator interface {
	GetMinIOEndpoint() string
	GetMinIOAccessKeyID() string
	GetMinIOSecretAccessKey() string
	GetMinIOSecure() bool
	GetMinIOInsecureSkipVerify() bool
	GetMinIOBucket() string
	GetMinIOBucketAutoCreate() bool
	GetMinIORegion() string
}

// Interface compatibility
var _ MinIOConfigurator = MinIO{}

// MinIO
type MinIO struct {
	MinIO *items.MinIO `mapstructure:"minio"`
}

// MinIONormalize
func (c MinIO) MinIONormalize() MinIO {
	if c.MinIO == nil {
		c.MinIO = items.NewMinIO()
	}
	return c
}

// GetMinIOEndpoint
func (c MinIO) GetMinIOEndpoint() string {
	return c.MinIO.GetEndpoint()
}

// GetMinIOAccessKeyID
func (c MinIO) GetMinIOAccessKeyID() string {
	return c.MinIO.GetAccessKeyID()
}

// GetMinIOSecretAccessKey
func (c MinIO) GetMinIOSecretAccessKey() string {
	return c.MinIO.GetSecretAccessKey()
}

// GetMinIOSecure
func (c MinIO) GetMinIOSecure() bool {
	return c.MinIO.GetSecure()
}

// GetMinIOInsecureSkipVerify
func (c MinIO) GetMinIOInsecureSkipVerify() bool {
	return c.MinIO.GetInsecureSkipVerify()
}

// GetMinIOBucket
func (c MinIO) GetMinIOBucket() string {
	return c.MinIO.GetBucket()
}

// GetMinIOBucketAutoCreate
func (c MinIO) GetMinIOBucketAutoCreate() bool {
	return c.MinIO.GetBucketAutoCreate()
}

// GetMinIORegion
func (c MinIO) GetMinIORegion() string {
	return c.MinIO.GetRegion()
}

// String
func (c MinIO) String() string {
	return fmt.Sprintf("MinIO=%s", c.MinIO)
}
