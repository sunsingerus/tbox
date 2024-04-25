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
	"os"
	"path/filepath"
)

// PathsConfigurator
type PathsConfigurator interface {
	GetAll(name string, opts ...*PathsOpts) []string
	Get(name string, opts ...*PathsOpts) string
	GetFile(filename, name string, opts ...*PathsOpts) string
}

// PathsOpts specifies paths options, such as base dir for rebase and fallback dir
type PathsOpts struct {
	// Base specifies the base on top of which to rebase relative paths.
	// In case base == nil no rebase required
	// In case *base == "" use CWD as a base
	// Otherwise rebase on top of *base, in case path is a relative one
	Base *string
	// Fallback specifies path which to fall back to in case specified paths name not found
	// In case fallback == nil no fallback required
	// In case *fallback == "" use CWD as a fallback
	// Otherwise fallback on *fallback
	Fallback *string
}

var (
	empty                = ""
	PathsOptsNothing     = &PathsOpts{}
	PathsOptsRebaseOnCWD = &PathsOpts{
		Base: &empty,
	}
	PathsOptsRebaseOnCWDFallbackOnCWD = &PathsOpts{
		Base:     &empty,
		Fallback: &empty,
	}
	PathsOptsDefault = PathsOptsNothing
)

// IMPORTANT
// IMPORTANT Do not forget to update String() function
// IMPORTANT

// Paths represents map key => one path
type Paths map[string]string

// MultiPaths represents map key => multiple paths
type MultiPaths map[string][]string

// Interface validation
var (
	_ PathsConfigurator = &Paths{}
	_ PathsConfigurator = &MultiPaths{}
)

// NewMultiPaths creates new MultiPaths
func NewMultiPaths() *MultiPaths {
	paths := new(MultiPaths)
	*paths = make(map[string][]string)
	return paths
}

// GetMap gets as map
func (f *MultiPaths) GetMap() map[string][]string {
	if f == nil {
		return nil
	}
	return *f
}

// GetNames gets list of names in paths
func (f *MultiPaths) GetNames() []string {
	if f == nil {
		return nil
	}
	var res []string
	for name := range *f {
		res = append(res, name)
	}
	return res
}

// getPaths is a getter
func (f *MultiPaths) getPaths(name string) []string {
	if f == nil {
		return nil
	}
	if *f == nil {
		return nil
	}
	if paths, ok := (*f)[name]; ok {
		return paths
	}
	return nil
}

// GetAll is a getter
func (f *MultiPaths) GetAll(name string, opts ...*PathsOpts) []string {
	var opt *PathsOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt == nil {
		opt = PathsOptsDefault
	}

	// Get paths by specified name
	paths := f.getPaths(name)

	// However, there may be no paths found by specified name,
	// so, let's check fallback options to fallback to in case no paths found
	if len(paths) < 1 {
		// There is no paths found. Need to fallback to possibly specified fallback paths
		switch {
		case opt.Fallback == nil:
			// No fallback path specified
		case *opt.Fallback == "":
			// Fallback to CWD
			if cwd, err := os.Getwd(); err == nil {
				// CWD found, all is fine
				paths = []string{
					cwd,
				}
			} else {
				// Unable to get CWD, fallback to root, instead of CWD
				paths = []string{
					"/",
				}
			}
		default:
			// Fallback to explicitly specified path
			paths = []string{
				*opt.Fallback,
			}
		}
	}

	// Variable "paths" should not be modified cause it points into somebody's mem
	// Make special result var to copy paths into
	var res []string

	// As we have possibly found paths, some of them may be relative and may require
	// to be rebased on top of some dir.
	// However, only relative paths should be rebased.
	for _, path := range paths {
		switch {
		case opt.Base == nil:
			// No rebase
			res = append(res, path)
		case *opt.Base == "":
			// Rebase on top CWD
			if filepath.IsAbs(path) {
				// Absolute path is not rebased and used as is
				res = append(res, path)
			} else {
				// Rebase relative path
				base, err := os.Getwd()
				if err != nil {
					// Unable to get CWD, fallback to root, instead of CWD
					base = "/"
				}
				res = append(res, filepath.Clean(filepath.Join(base, path)))
			}
		default:
			// Rebase on top of explicitly specified path
			if filepath.IsAbs(path) {
				// Absolute path is not rebased and used as is
				res = append(res, path)
			} else {
				base := *opt.Base
				res = append(res, filepath.Clean(filepath.Join(base, path)))
			}
		}
	}

	return res
}

// Get gets the first path
func (f *MultiPaths) Get(name string, opts ...*PathsOpts) string {
	paths := f.GetAll(name, opts...)
	if len(paths) > 0 {
		return paths[0]
	}
	return ""
}

// GetFile gets a file based on paths
func (f *MultiPaths) GetFile(filename, name string, opts ...*PathsOpts) string {
	if path := f.Get(name, opts...); path == "" {
		return filename
	} else {
		return filepath.Join(path, filename)
	}
}

// String is a stringifier
func (f *MultiPaths) String() string {
	if f == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "%v\n", *f)

	return b.String()
}

// NewPaths creates new Paths
func NewPaths() *Paths {
	paths := new(Paths)
	*paths = make(map[string]string)
	return paths
}

// GetMap gets as map
func (f *Paths) GetMap() map[string]string {
	if f == nil {
		return nil
	}
	return *f
}

// GetNames gets list of names in paths
func (f *Paths) GetNames() []string {
	if f == nil {
		return nil
	}
	var res []string
	for name := range *f {
		res = append(res, name)
	}
	return res
}

// getPath is a getter
func (f *Paths) getPath(name string) string {
	if f == nil {
		return ""
	}
	if *f == nil {
		return ""
	}
	if paths, ok := (*f)[name]; ok {
		return paths
	}
	return ""
}

// GetAll is a getter
func (f *Paths) GetAll(name string, opts ...*PathsOpts) []string {
	var opt *PathsOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt == nil {
		opt = PathsOptsDefault
	}

	// Get paths by specified name
	var paths []string

	if f.getPath(name) != "" {
		paths = []string{f.getPath(name)}
	}

	// However, there may be no paths found by specified name,
	// so, let's check fallback options to fallback to in case no paths found
	if len(paths) < 1 {
		// There is no paths found. Need to fallback to possibly specified fallback paths
		switch {
		case opt.Fallback == nil:
			// No fallback path specified
		case *opt.Fallback == "":
			// Fallback to CWD
			if cwd, err := os.Getwd(); err == nil {
				// CWD found, all is fine
				paths = []string{
					cwd,
				}
			} else {
				// Unable to get CWD, fallback to root, instead of CWD
				paths = []string{
					"/",
				}
			}
		default:
			// Fallback to explicitly specified path
			paths = []string{
				*opt.Fallback,
			}
		}
	}

	// Variable "paths" should not be modified cause it points into somebody's mem
	// Make special result var to copy paths into
	var res []string

	// As we have possibly found paths, some of them may be relative and may require
	// to be rebased on top of some dir.
	// However, only relative paths should be rebased.
	for _, path := range paths {
		switch {
		case opt.Base == nil:
			// No rebase
			res = append(res, path)
		case *opt.Base == "":
			// Rebase on top CWD
			if filepath.IsAbs(path) {
				// Absolute path is not rebased and used as is
				res = append(res, path)
			} else {
				// Rebase relative path
				base, err := os.Getwd()
				if err != nil {
					// Unable to get CWD, fallback to root, instead of CWD
					base = "/"
				}
				res = append(res, filepath.Clean(filepath.Join(base, path)))
			}
		default:
			// Rebase on top of explicitly specified path
			if filepath.IsAbs(path) {
				// Absolute path is not rebased and used as is
				res = append(res, path)
			} else {
				base := *opt.Base
				res = append(res, filepath.Clean(filepath.Join(base, path)))
			}
		}
	}

	return res
}

// Get gets the first path
func (f *Paths) Get(name string, opts ...*PathsOpts) string {
	paths := f.GetAll(name, opts...)
	if len(paths) > 0 {
		return paths[0]
	}
	return ""
}

// GetFile gets a file based on paths
func (f *Paths) GetFile(filename, name string, opts ...*PathsOpts) string {
	if path := f.Get(name, opts...); path == "" {
		return filename
	} else {
		return filepath.Join(path, filename)
	}
}

// String is a stringifier
func (f *Paths) String() string {
	if f == nil {
		return nilString
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "%v\n", *f)

	return b.String()
}
