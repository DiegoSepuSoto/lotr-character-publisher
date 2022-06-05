package mapper

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
	"github.com/segmentio/kafka-go"

	"DiegoSepuSoto/lotr-character-publisher/src/domain/models"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func BuildKafkaMessages(characters []*models.Character) []kafka.Message {
	return lo.Map(characters, func(character *models.Character, i int) kafka.Message {
		characterJSON, _ := json.Marshal(character)

		return kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: characterJSON,
		}
	})
}
