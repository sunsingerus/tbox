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

package devcerts

import (
	"path/filepath"
	"runtime"
)

// dir is the directory of this package.
var dir string

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	dir = filepath.Dir(currentFile)
}

// Path returns the absolute path for the given relative file.
// Path returned is relative to the directory where source code of this module resides.
// If provided path is already absolute, it is returned unmodified.
func Path(paths ...string) string {
	path := filepath.Join(paths...)

	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(dir, path)
}
