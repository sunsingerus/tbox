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

const (
	ReportTypeReserved    = 0
	ReportTypeUnspecified = 100
)

var ReportTypeEnum = NewEnum()

func init() {
	ReportTypeEnum.MustRegister("ReportTypeReserved", ReportTypeReserved)
	ReportTypeEnum.MustRegister("ReportTypeUnspecified", ReportTypeUnspecified)
}

// NewReport
func NewReport() *Report {
	return new(Report)
}

// EnsureHeader
func (x *Report) EnsureHeader() *Metadata {
	if x == nil {
		return nil
	}
	x.Header = new(Metadata)
	return x.Header
}

// SetBytes
func (x *Report) SetBytes(bytes []byte) *Report {
	if x == nil {
		return nil
	}
	x.Bytes = bytes
	return x
}

// HasSubReports
func (x *Report) HasSubReports() bool {
	if x == nil {
		return false
	}
	return len(x.Children) > 0
}

// AddSubReport
func (x *Report) AddSubReport(r *Report) *Report {
	if x == nil {
		return nil
	}
	x.Children = append(x.Children, r)
	return x
}

// WalkSubReports
func (x *Report) WalkSubReports(f func(i int, subReport *Report)) *Report {
	if x == nil {
		return nil
	}
	for i, subReport := range x.Children {
		f(i, subReport)
	}
	return x
}

// String
func (x *Report) String() string {
	return "to be implemented"
}
