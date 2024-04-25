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

package text

import (
	"os"
	"unicode/utf8"
)

// isText is an internal function to check whether provided bytes are a text.
// For being a text it should be
// 1. either a string or a []byte
// 2. contain valid letters or numbers
func isText(s []byte) bool {
	// Make a decision based on 'num' bytes at max
	// Prepare buf with 'num' bytes at max
	const num = 1024
	if len(s) > num {
		s = s[0:num]
	}

	for i, c := range string(s) {
		if i+utf8.UTFMax > len(s) {
			// Last char may be incomplete - ignore
			break
		}

		if c == 0xFFFD {
			// Unicode Character 'REPLACEMENT CHARACTER'
			// used to replace an incoming character whose value is unknown or unrepresentable in Unicode
			// This is binary data.
			return false
		}

		if (c < ' ') && // Non-printable char
			(c != '\n') && // Newline
			(c != '\t') && // Tab
			(c != '\f') { // Form feed
			// Non-printable and non-control char.
			// This is binary data
			return false
		}
	}

	// This is a text
	return true
}

// IsText checks whether provided variable is a text.
// Accepts string or []byte
func IsText(a interface{}) bool {
	switch typed := a.(type) {
	case string:
		return isText([]byte(typed))
	case []byte:
		return isText(typed)
	}
	return false
}

// IsTextFile check whether provided filename specifies path to a readable file with valid text content
func IsTextFile(filename string) (bool, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	return IsText(data), nil
}
