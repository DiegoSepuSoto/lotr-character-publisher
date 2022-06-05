package main

import (
	"DiegoSepuSoto/lotr-character-publisher/src/application/usecase/lotr/information"
	lotrJSONRepository "DiegoSepuSoto/lotr-character-publisher/src/infrastructure/data/json/lotr"
	"DiegoSepuSoto/lotr-character-publisher/src/infrastructure/events/kafka/publisher"
	lotrKafkaPublisherRepository "DiegoSepuSoto/lotr-character-publisher/src/infrastructure/events/kafka/publisher/lotr"
	"DiegoSepuSoto/lotr-character-publisher/src/utils/logger"
	"os"
)

func main() {
	log := logger.GetLogger("application", "main")
	lotrJSONRepositoryImpl := lotrJSONRepository.NewLOTRJSONRepository()
	kafkaWriter := publisher.NewKafkaWriter(os.Getenv("KAFKA_BROKER"), os.Getenv("KAFKA_TOPIC"))
	publisherRepository := lotrKafkaPublisherRepository.NewLOTRKafkaPublisherRepository(kafkaWriter)

	informationUseCase := information.NewLOTRInformationUseCase(lotrJSONRepositoryImpl, publisherRepository)

	err := informationUseCase.PublishCharactersInformation("data.json")
	if err != nil {
		log.Errorln("failed to get characters")
	}
}
