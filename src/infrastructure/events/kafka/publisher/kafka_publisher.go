package publisher

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter interface {
	WriteMessages(ctx context.Context, messages ...kafka.Message) error
}

func NewKafkaWriter(address, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(address),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
