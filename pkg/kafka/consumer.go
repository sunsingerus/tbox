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
	log "github.com/sirupsen/logrus"

	"github.com/sunsingerus/tbox/pkg/api/common"
	"github.com/sunsingerus/tbox/pkg/config/sections"
	"github.com/sunsingerus/tbox/pkg/softwareid"
)

// Consumer
type Consumer struct {
	endpoint *common.KafkaEndpoint
	address  *common.KafkaAddress

	config            *sarama.Config
	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer
}

// NewConsumer
func NewConsumer(endpoint *common.KafkaEndpoint, address *common.KafkaAddress) *Consumer {
	var err error

	c := &Consumer{}
	c.endpoint = endpoint
	c.address = address
	c.config = sarama.NewConfig()
	c.config.Version = sarama.V2_0_0_0
	c.config.ClientID = softwareid.Name
	c.consumer, err = sarama.NewConsumer(c.endpoint.Brokers, c.config)
	if err != nil {
		c.Close()
		log.Errorf("unable to create new Kafka consumer. err: %v", err)
		return nil
	}

	topics, err := c.consumer.Topics()
	if err != nil {
		c.Close()
		log.Errorf("unable to list topics. err: %v", err)
		return nil
	}

	partitions, err := c.consumer.Partitions(c.address.Topic)
	if err != nil {
		c.Close()
		log.Errorf("unable to list partitions. err: %v", err)
		return nil
	}

	log.Info("Going to consume:")
	log.Infof("topic %s of %v", c.address.Topic, topics)
	log.Infof("partition %d of %v", c.address.Partition, partitions)

	c.partitionConsumer, err = c.consumer.ConsumePartition(c.address.Topic, c.address.Partition, sarama.OffsetNewest)
	if err != nil {
		c.Close()
		log.Errorf("unable to consume partition. err: %v", err)
		return nil
	}

	return c
}

// NewConsumerConfig
func NewConsumerConfig(cfg sections.KafkaConfigurator, topic string) *Consumer {
	return NewConsumer(cfg.GetKafkaEndpoint(), common.NewKafkaAddress(topic, 0))
}

// Close will close partition consumer and drain partition consumer's Messages() chan, so blocking Messages() will exit
func (c *Consumer) Close() {
	if c.partitionConsumer != nil {
		_ = c.partitionConsumer.Close()
		c.partitionConsumer = nil
	}

	if c.consumer != nil {
		_ = c.consumer.Close()
		c.consumer = nil
	}
}

// Recv is a blocking call
func (c *Consumer) Recv() *sarama.ConsumerMessage {
	msg := <-c.partitionConsumer.Messages()
	if msg != nil {
		log.Infof("Got message %s", MsgAddressPrintable(msg))
	}
	return msg
}
