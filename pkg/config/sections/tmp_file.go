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

// TmpFileConfigurator
type TmpFileConfigurator interface {
	GetTmpFileDir() string
	GetTmpFilePattern() string
}

// Interface compatibility
var _ TmpFileConfigurator = TmpFile{}

// TmpFile
type TmpFile struct {
	TmpFile *items.DirFile `mapstructure:"tmpfile"`
}

// TmpFileNormalize
func (c TmpFile) TmpFileNormalize() TmpFile {
	if c.TmpFile == nil {
		c.TmpFile = items.NewDirFile()
	}
	return c
}

// GetTmpFileDir
func (c TmpFile) GetTmpFileDir() string {
	return c.TmpFile.GetDir()
}

// GetTmpFilePattern
func (c TmpFile) GetTmpFilePattern() string {
	return c.TmpFile.GetPattern()
}

// String
func (c TmpFile) String() string {
	return fmt.Sprintf("TmpFile=%s", c.TmpFile)
}
