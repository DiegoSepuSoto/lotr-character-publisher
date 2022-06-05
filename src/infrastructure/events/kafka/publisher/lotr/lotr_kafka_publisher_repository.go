package lotr

import "DiegoSepuSoto/lotr-character-publisher/src/infrastructure/events/kafka/publisher"

type lotrPublisherRepository struct {
	kafkaPublisher publisher.KafkaWriter
}

func NewLOTRKafkaPublisherRepository(kafkaPublisher publisher.KafkaWriter) *lotrPublisherRepository {
	return &lotrPublisherRepository{kafkaPublisher: kafkaPublisher}
}
