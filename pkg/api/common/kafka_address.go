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

import "fmt"

// NewKafkaAddress
func NewKafkaAddress(topic string, partition int32) *KafkaAddress {
	return &KafkaAddress{
		Topic:     topic,
		Partition: partition,
	}
}

// GetTopics
func (x *KafkaAddress) GetTopics() []string {
	if x != nil {
		return []string{x.Topic}
	}

	return nil
}

// String
func (x *KafkaAddress) String() string {
	if x != nil {
		return fmt.Sprintf("%s/%d", x.Topic, x.Partition)
	}
	return ""
}
