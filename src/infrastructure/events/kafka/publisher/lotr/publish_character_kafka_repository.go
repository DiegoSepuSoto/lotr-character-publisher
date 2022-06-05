package lotr

import (
	"context"

	"DiegoSepuSoto/lotr-character-publisher/src/domain/models"
	"DiegoSepuSoto/lotr-character-publisher/src/infrastructure/events/kafka/publisher/lotr/mapper"
	"DiegoSepuSoto/lotr-character-publisher/src/utils/logger"
)

func (r *lotrPublisherRepository) PublishCharacters(characters []*models.Character) error {
	log := logger.GetLogger("repository", "PublishCharacter")
	kafkaMessages := mapper.BuildKafkaMessages(characters)

	err := r.kafkaPublisher.WriteMessages(context.Background(), kafkaMessages...)
	if err != nil {
		log.Errorf("couldn't write message to kafka: [%s]", err.Error())
		return err
	}

	log.Errorln("messages written successfully")

	return nil
}
