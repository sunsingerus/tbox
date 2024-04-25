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

package json

import (
	j "encoding/json"
	"os"
)

// IsJSON checks whether provided variable is a JSON. For being a JSON it should be
// 1. either a string or a []byte
// 2. contain valid JSON
func IsJSON(a interface{}) bool {
	var js j.RawMessage
	switch typed := a.(type) {
	case string:
		return j.Unmarshal([]byte(typed), &js) == nil
	case []byte:
		return j.Unmarshal(typed, &js) == nil
	}
	return false
}

// IsJSONFile check whether filename specifies path to a readable file with valid JSON content
func IsJSONFile(filename string) (bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	return IsJSON(data), nil
}
