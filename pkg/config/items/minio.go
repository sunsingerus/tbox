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

package items

import (
	"bytes"
	"fmt"
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT
type MinIO struct {
	Enabled bool `mapstructure:"enabled"`
	// MinIO connection
	Endpoint           string `mapstructure:"endpoint"`
	AccessKeyID        string `mapstructure:"accessKeyID"`
	SecretAccessKey    string `mapstructure:"secretAccessKey"`
	Secure             bool   `mapstructure:"secure"`
	InsecureSkipVerify bool   `mapstructure:"insecureSkipVerify"`
	// MinIO internals
	Bucket           string `mapstructure:"bucket"`
	BucketAutoCreate bool   `mapstructure:"bucketAutoCreate"`
	Region           string `mapstructure:"region"`
	// IMPORTANT
	// IMPORTANT Do not forget to update String() function
	// IMPORTANT
}

// NewMinIO is a constructor
func NewMinIO() *MinIO {
	return new(MinIO)
}

// GetEnabled is a getter
func (m *MinIO) GetEnabled() bool {
	if m == nil {
		return false
	}
	return m.Enabled
}

// GetEndpoint is a getter
func (m *MinIO) GetEndpoint() string {
	if m == nil {
		return ""
	}
	return m.Endpoint
}

// GetAccessKeyID is a getter
func (m *MinIO) GetAccessKeyID() string {
	if m == nil {
		return ""
	}
	return m.AccessKeyID
}

// GetSecretAccessKey is a getter
func (m *MinIO) GetSecretAccessKey() string {
	if m == nil {
		return ""
	}
	return m.SecretAccessKey
}

// GetSecure is a getter
func (m *MinIO) GetSecure() bool {
	if m == nil {
		return false
	}
	return m.Secure
}

// GetInsecureSkipVerify is a getter
func (m *MinIO) GetInsecureSkipVerify() bool {
	if m == nil {
		return false
	}
	return m.InsecureSkipVerify
}

// GetBucket is a getter
func (m *MinIO) GetBucket() string {
	if m == nil {
		return ""
	}
	return m.Bucket
}

// GetBucketAutoCreate is a getter
func (m *MinIO) GetBucketAutoCreate() bool {
	if m == nil {
		return false
	}
	return m.BucketAutoCreate
}

// GetRegion is a getter
func (m *MinIO) GetRegion() string {
	if m == nil {
		return ""
	}
	return m.Region
}

// String is a stringifier
func (m *MinIO) String() string {
	if m == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Enabled: %v\n", m.Enabled)
	_, _ = fmt.Fprintf(b, "EndpointID: %v\n", m.Endpoint)
	_, _ = fmt.Fprintf(b, "AccessKeyID: %v\n", m.AccessKeyID)
	_, _ = fmt.Fprintf(b, "SecretAccessKey: %v\n", m.SecretAccessKey)
	_, _ = fmt.Fprintf(b, "Secure: %v\n", m.Secure)
	_, _ = fmt.Fprintf(b, "InsecureSkipVerify: %v\n", m.InsecureSkipVerify)
	_, _ = fmt.Fprintf(b, "Bucket: %v\n", m.Bucket)
	_, _ = fmt.Fprintf(b, "BucketAutoCreate: %v\n", m.BucketAutoCreate)
	_, _ = fmt.Fprintf(b, "Region: %v\n", m.Region)

	return b.String()
}
