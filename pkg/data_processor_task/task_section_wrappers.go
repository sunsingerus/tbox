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

package data_processor_task

import (
	"path/filepath"
	"strconv"

	"github.com/sunsingerus/tbox/pkg/util"
)

/*******************************************/
/*******************************************/
/*                                         */
/*             Section Wrappers            */
/*                                         */
/*******************************************/
/*******************************************/

// GetConfigFiles gets all config files
func (t *DataProcessorTask) GetConfigFiles() []string {
	return t.GetAll(ConfigFiles)
}

// HasConfigFiles checks whether there are config file(s)
func (t *DataProcessorTask) HasConfigFiles() bool {
	return t.Has(ConfigFiles)
}

// GetConfigFile gets the first config file
func (t *DataProcessorTask) GetConfigFile(defaultValue ...string) string {
	return t.Get(ConfigFiles, defaultValue...)
}

// SetConfigFile adds config file(s)
func (t *DataProcessorTask) SetConfigFile(file ...string) *DataProcessorTask {
	return t.Set(ConfigFiles, file...)
}

// AddConfigFile adds config file(s)
func (t *DataProcessorTask) AddConfigFile(file ...string) *DataProcessorTask {
	return t.Add(ConfigFiles, file...)
}

// GetConfigDirs gets all config dirs
func (t *DataProcessorTask) GetConfigDirs() []string {
	return t.GetAll(ConfigDirs)
}

// HasConfigDirs checks whether there are config dirs(s)
func (t *DataProcessorTask) HasConfigDirs() bool {
	return t.Has(ConfigDirs)
}

// GetConfigDir gets the first config dir
func (t *DataProcessorTask) GetConfigDir(defaultValue ...string) string {
	return t.Get(ConfigDirs, defaultValue...)
}

// SetConfigDir sets config dir(s)
func (t *DataProcessorTask) SetConfigDir(dir ...string) *DataProcessorTask {
	return t.Set(ConfigDirs, dir...)
}

// AddConfigDir adds config dir(s)
func (t *DataProcessorTask) AddConfigDir(dir ...string) *DataProcessorTask {
	return t.Add(ConfigDirs, dir...)
}

// GetInputFiles gets all input files
func (t *DataProcessorTask) GetInputFiles() []string {
	return t.GetAll(InputFiles)
}

// HasInputFiles checks whether there are input file(s)
func (t *DataProcessorTask) HasInputFiles() bool {
	return t.Has(InputFiles)
}

// GetInputFileByIndex gets input file with specified index
func (t *DataProcessorTask) GetInputFileByIndex(index int, defaultValue ...string) string {
	return t.GetByIndex(InputFiles, index, defaultValue...)
}

// GetInputFilePathByIndex gets input file with specified index with full path
func (t *DataProcessorTask) GetInputFilePathByIndex(index int, defaultValue ...string) string {
	file, found := t.FetchByIndex(InputFiles, index, defaultValue...)
	if !found {
		// Nothing found, let's return default value
		return file
	}

	return t.GetInputPath(file)
}

// GetInputFile gets the first input file
func (t *DataProcessorTask) GetInputFile(defaultValue ...string) string {
	return t.Get(InputFiles, defaultValue...)
}

// GetInputFilePath gets the first input file
func (t *DataProcessorTask) GetInputFilePath(defaultValue ...string) string {
	file, found := t.Fetch(InputFiles, defaultValue...)
	if !found {
		// Nothing found, let's return default value
		return file
	}

	return t.GetInputPath(file)
}

// GetInputPath makes input path for specified part
func (t *DataProcessorTask) GetInputPath(input string) string {
	// Test filename itself
	path := filepath.Join(input)
	if filepath.IsAbs(path) {
		return path
	}

	// Test InputDir()/filename
	path = filepath.Join(t.GetInputDir(), input)
	if filepath.IsAbs(path) {
		return path
	}

	// RootDir()/InputDir()/filename
	path = filepath.Join(t.GetRootDir(), t.GetInputDir(), input)
	return path
}

// SetInputFile sets input file(s)
func (t *DataProcessorTask) SetInputFile(file ...string) *DataProcessorTask {
	return t.Set(InputFiles, file...)
}

// AddInputFile adds input file(s)
func (t *DataProcessorTask) AddInputFile(file ...string) *DataProcessorTask {
	return t.Add(InputFiles, file...)
}

// GetInputDirs gets all input dirs
func (t *DataProcessorTask) GetInputDirs() []string {
	return t.GetAll(InputDirs)
}

// HasInputDirs checks whether there are input dires(s)
func (t *DataProcessorTask) HasInputDirs() bool {
	return t.Has(InputDirs)
}

// GetInputDir gets the first input dir
func (t *DataProcessorTask) GetInputDir(defaultValue ...string) string {
	return t.Get(InputDirs, defaultValue...)
}

// SetInputDir sets input dir(s)
func (t *DataProcessorTask) SetInputDir(dir ...string) *DataProcessorTask {
	return t.Set(InputDirs, dir...)
}

// AddInputDir adds input dir(s)
func (t *DataProcessorTask) AddInputDir(dir ...string) *DataProcessorTask {
	return t.Add(InputDirs, dir...)
}

// GetOutputFiles gets all output files
func (t *DataProcessorTask) GetOutputFiles() []string {
	return t.GetAll(OutputFiles)
}

// HasOutputFiles checks whether there are output file(s)
func (t *DataProcessorTask) HasOutputFiles() bool {
	return t.Has(OutputFiles)
}

// GetOutputFile gets the first output file
func (t *DataProcessorTask) GetOutputFile(defaultValue ...string) string {
	return t.Get(OutputFiles, defaultValue...)
}

// GetOutputFilePath gets the first output file
func (t *DataProcessorTask) GetOutputFilePath(defaultValue ...string) string {
	file, found := t.Fetch(OutputFiles, defaultValue...)
	if !found {
		// Nothing found, let's return default value
		return file
	}

	return t.GetOutputPath(file)
}

// GetOutputPath makes output path for specified part
func (t *DataProcessorTask) GetOutputPath(output string) string {
	// Test filename itself
	path := filepath.Join(output)
	if filepath.IsAbs(path) {
		return path
	}

	// Test OutputDir()/filename
	path = filepath.Join(t.GetOutputDir(), output)
	if filepath.IsAbs(path) {
		return path
	}

	// RootDir()/OutputDir()/filename
	path = filepath.Join(t.GetRootDir(), t.GetOutputDir(), output)
	return path
}

// SetOutputFile sets output file(s)
func (t *DataProcessorTask) SetOutputFile(file ...string) *DataProcessorTask {
	return t.Set(OutputFiles, file...)
}

// AddOutputFile adds output file(s)
func (t *DataProcessorTask) AddOutputFile(file ...string) *DataProcessorTask {
	return t.Add(OutputFiles, file...)
}

// GetOutputDirs gets all output dirs
func (t *DataProcessorTask) GetOutputDirs() []string {
	return t.GetAll(OutputDirs)
}

// HasOutputDirs checks whether there are output dir(s)
func (t *DataProcessorTask) HasOutputDirs() bool {
	return t.Has(OutputDirs)
}

// GetOutputDir gets the first output dir
func (t *DataProcessorTask) GetOutputDir(defaultValue ...string) string {
	return t.Get(OutputDirs, defaultValue...)
}

// SetOutputDir sets output dir(s)
func (t *DataProcessorTask) SetOutputDir(dir ...string) *DataProcessorTask {
	return t.Set(OutputDirs, dir...)
}

// AddOutputDir adds output dir(s)
func (t *DataProcessorTask) AddOutputDir(dir ...string) *DataProcessorTask {
	return t.Add(OutputDirs, dir...)
}

// GetInputTables gets all input tables
func (t *DataProcessorTask) GetInputTables() []string {
	return t.GetAll(InputTables)
}

// HasInputTables checks whether there are input table(s)
func (t *DataProcessorTask) HasInputTables() bool {
	return t.Has(InputTables)
}

// GetInputTable gets the first input table
func (t *DataProcessorTask) GetInputTable(defaultValue ...string) string {
	return t.Get(InputTables, defaultValue...)
}

// SetInputTable sets input table(s)
func (t *DataProcessorTask) SetInputTable(table ...string) *DataProcessorTask {
	return t.Set(InputTables, table...)
}

// AddInputTable adds input table(s)
func (t *DataProcessorTask) AddInputTable(table ...string) *DataProcessorTask {
	return t.Add(InputTables, table...)
}

// GetOutputTables gets all output tables
func (t *DataProcessorTask) GetOutputTables() []string {
	return t.GetAll(OutputTables)
}

// HasOutputTables checks whether there are output table(s)
func (t *DataProcessorTask) HasOutputTables() bool {
	return t.Has(OutputTables)
}

// GetOutputTable gets the first output table
func (t *DataProcessorTask) GetOutputTable(defaultValue ...string) string {
	return t.Get(OutputTables, defaultValue...)
}

// SetOutputTable sets output table(s)
func (t *DataProcessorTask) SetOutputTable(table ...string) *DataProcessorTask {
	return t.Set(OutputTables, table...)
}

// AddOutputTable adds output table(s)
func (t *DataProcessorTask) AddOutputTable(table ...string) *DataProcessorTask {
	return t.Add(OutputTables, table...)
}

// HasReportLevel checks presence
func (t *DataProcessorTask) HasReportLevel() bool {
	return t.Len(ReportLevel) > 0
}

// SetReportLevel sets
func (t *DataProcessorTask) SetReportLevel(level int) *DataProcessorTask {
	return t.Set(ReportLevel, strconv.Itoa(level))
}

// HasSummaryLevel checks presence
func (t *DataProcessorTask) HasSummaryLevel() bool {
	return t.Len(SummaryLevel) > 0
}

// SetSummaryLevel sets
func (t *DataProcessorTask) SetSummaryLevel(level int) *DataProcessorTask {
	return t.Set(SummaryLevel, strconv.Itoa(level))
}

// HasTraceLevel checks presence
func (t *DataProcessorTask) HasTraceLevel() bool {
	return t.Len(TraceLevel) > 0
}

// SetTraceLevel sets
func (t *DataProcessorTask) SetTraceLevel(level int) *DataProcessorTask {
	return t.Set(TraceLevel, strconv.Itoa(level))
}

// WalkInputFiles walks with the specified callable over all input files.
// provideFullFilepath specifies whether callable will receive full filepath or just a filename
func (t *DataProcessorTask) WalkInputFiles(provideFullFilepath bool, f func(string) bool) *DataProcessorTask {
	return t.WalkSection(InputFiles, func(item string) bool {
		if provideFullFilepath {
			item = util.FullPath(t.GetRootDir(), t.GetInputDir(), item)
		}
		return f(item)
	})
}

// WalkOutputFiles walks with the specified callable over all output files.
// provideFullFilepath specifies whether callable will receive full filepath or just a filename
func (t *DataProcessorTask) WalkOutputFiles(provideFullFilepath bool, f func(string) bool) *DataProcessorTask {
	return t.WalkSection(OutputFiles, func(item string) bool {
		if provideFullFilepath {
			item = util.FullPath(t.GetRootDir(), t.GetOutputDir(), item)
		}
		return f(item)
	})
}
