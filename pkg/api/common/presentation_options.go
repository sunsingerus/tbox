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

import log "github.com/sirupsen/logrus"

// NewPresentationOptions
func NewPresentationOptions() *PresentationOptions {
	return new(PresentationOptions)
}

// HasEncoding
func (x *PresentationOptions) HasEncoding() bool {
	if x == nil {
		return false
	}
	return x.Encoding != nil
}

// SetEncoding
func (x *PresentationOptions) SetEncoding(encoding *Encoding) *PresentationOptions {
	if x == nil {
		return nil
	}
	x.Encoding = encoding
	return x
}

// HasCompression
func (x *PresentationOptions) HasCompression() bool {
	if x == nil {
		return false
	}
	return x.Compression != nil
}

// SetCompression
func (x *PresentationOptions) SetCompression(compression *Compression) *PresentationOptions {
	if x == nil {
		return nil
	}
	x.Compression = compression
	return x
}

// String
func (x *PresentationOptions) String() string {
	return "to be implemented"
}

// Log
func (x *PresentationOptions) Log() {
	if x == nil {
		return
	}
	log.Infof("presentation options: %s", x)
}
