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

package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/softwareid"
	log "github.com/sirupsen/logrus"
)

// Producer
type Producer struct {
	endpoint *common.KafkaEndpoint
	address  *common.KafkaAddress

	config   *sarama.Config
	producer sarama.SyncProducer
}

// NewProducer
func NewProducer(endpoint *common.KafkaEndpoint, address *common.KafkaAddress) *Producer {
	var err error

	p := &Producer{}
	p.endpoint = endpoint
	p.address = address
	p.config = sarama.NewConfig()
	p.config.Version = sarama.V2_0_0_0
	p.config.ClientID = softwareid.Name
	// If this config is used to create a `SyncProducer`, both must be set to true,
	// and you shall not read from the channels since the producer does this internally.
	p.config.Producer.Return.Successes = true
	p.config.Producer.Return.Errors = true
	p.producer, err = sarama.NewSyncProducer(p.endpoint.GetBrokers(), p.config)
	if err != nil {
		log.Errorf("unable to create NewSyncProducer(brokers:%v). err: %v", p.endpoint.GetBrokers(), err)
		p.Close()
		return nil
	}

	return p
}

// NewProducerConfig
func NewProducerConfig(cfg sections.KafkaConfigurator) *Producer {
	return NewProducer(
		cfg.GetKafkaEndpoint(),
		common.NewKafkaAddress(cfg.GetKafkaTopic(), 0),
	)
}

// SetAddress
func (p *Producer) SetAddress(address *common.KafkaAddress) *Producer {
	p.address = address
	return p
}

// SetTopic
func (p *Producer) SetTopic(topic string) *Producer {
	p.address = common.NewKafkaAddress(topic, 0)
	return p
}

// GetTopic
func (p *Producer) GetTopic() string {
	return p.address.GetTopic()
}

// CreateTopic
func (p *Producer) CreateTopic() (error, error) {
	var er2 error

	admin, err := sarama.NewClusterAdmin(p.endpoint.GetBrokers(), p.config)
	if err != nil {
		return err, er2
	}

	detail := &sarama.TopicDetail{
		NumPartitions:     int32(1),
		ReplicationFactor: int16(1),
	}

	err = admin.CreateTopic(p.GetTopic(), detail, false)
	if err != nil {
		if t, ok := err.(*sarama.TopicError); ok {
			if t.Err == sarama.ErrTopicAlreadyExists {
				// In case topic already exists, do not treat it as an error
				er2 = err
				err = nil
			}
		}
	}
	admin.Close()
	return err, er2
}

// ListTopics
func (p *Producer) ListTopics() (map[string]sarama.TopicDetail, error) {
	admin, err := sarama.NewClusterAdmin(p.endpoint.GetBrokers(), p.config)
	if err != nil {
		return nil, err
	}

	_map, err := admin.ListTopics()
	admin.Close()
	return _map, err
}

// Close
func (p *Producer) Close() {
	if p.producer != nil {
		_ = p.producer.Close()
		p.producer = nil
	}
}

// Send
func (p *Producer) Send(data []byte) error {

	msg := &sarama.ProducerMessage{
		Topic: p.address.Topic,
		Value: sarama.ByteEncoder(data),
		// Key
		// Headers - relayed to consumer
		// Metadata - relayed to the Successes and Errors channels
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Errorf("FAILED to send message: %s", err)
	} else {
		log.Infof("message sent to %s", MsgAddressPrintable(msg))
	}

	return err
}
