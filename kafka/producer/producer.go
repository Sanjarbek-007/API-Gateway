package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type ProducerIkafka interface {
	Producermessage(topic string, msg []byte) error
	Close()
}

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducerInit(brokers []string) (*KafkaProducer, error) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...),
		AllowAutoTopicCreation: true,
	}

	return &KafkaProducer{writer: writer}, nil

}

func (k *KafkaProducer) Producermessage(topic string, msg []byte) error {
	return k.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: msg,
	})
}

func (k *KafkaProducer) WriteToNotification(topic string, msg []byte) error {
	return k.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: msg,
	})
}


func (k *KafkaProducer) Close() {
	k.writer.Close()
}
