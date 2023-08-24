package lib

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

var subConn sarama.Consumer
var pubConn sarama.SyncProducer

type Event struct {
	Event string `json:"event"`
	ID    string `json:"ID"`
	Param string `json:"param"`
}

func KafkaPubConn() (sarama.SyncProducer, error) {
	if pubConn == nil {
		// brokersUrl := url
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 5
		var err error
		pubConn, err = sarama.NewSyncProducer([]string{"kafka:9092"}, config)
		if err != nil {
			return nil, err
		}
	}
	return pubConn, nil
}

func KafkaSubConn() (sarama.Consumer, error) {
	if subConn == nil {
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		var err error
		subConn, err = sarama.NewConsumer([]string{"kafka:9092"}, config)
		if err != nil {
			return nil, err
		}
	}
	return subConn, nil
}

func PubKafka(topic string, event Event) error {
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err = pubConn.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func SubKafka(topic string, c chan *sarama.ConsumerMessage, e chan sarama.ConsumerError, offset int64) {
	// Subscribe to connection
	consumer, err := subConn.ConsumePartition(topic, int32(0), offset)
	if err != nil {
		e <- sarama.ConsumerError{
			Topic:     topic,
			Partition: 0,
			Err:       err,
		}
		return
	}

	for {
		msg := <-consumer.Messages()
		c <- msg
	}
}
