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

// NewObjectsRequest is a constructor
func NewObjectsRequest() *ObjectsRequest {
	return new(ObjectsRequest)
}

// SetRequestDomain is a setter
func (x *ObjectsRequest) SetRequestDomain(domain *Domain) *ObjectsRequest {
	if x == nil {
		return nil
	}
	x.RequestDomain = domain
	return x
}

// SetResultDomain is a setter
func (x *ObjectsRequest) SetResultDomain(domain *Domain) *ObjectsRequest {
	if x == nil {
		return nil
	}
	x.ResultDomain = domain
	return x
}

// GetRequestsNum gets number of requests
func (x *ObjectsRequest) GetRequestsNum() int {
	return len(x.GetRequests())
}

// Append appends request(s)
func (x *ObjectsRequest) Append(request ...*ObjectRequest) *ObjectsRequest {
	if x == nil {
		return nil
	}
	x.Requests = append(x.Requests, request...)
	return x
}

// String is a stringifier
func (x *ObjectsRequest) String() string {
	return "to be implemented"
}
