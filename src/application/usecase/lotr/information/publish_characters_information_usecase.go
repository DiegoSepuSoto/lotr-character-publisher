package information

import (
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"

	"DiegoSepuSoto/lotr-character-publisher/src/domain/models"
	"DiegoSepuSoto/lotr-character-publisher/src/utils/logger"
)

func (u *lotrInformationUseCase) PublishCharactersInformation(filePath string) error {
	log := logger.GetLogger("useCase", "PublishCharactersInformation")

	characters, err := u.lotrInformation.GetCharactersInformation(filePath)
	if err != nil {
		return err
	}
	filteredCharacters := deleteDuplicatedCharacters(log, characters)

	log.Printf("%d duplicated characters found", len(characters)-len(filteredCharacters))

	err = u.lotrPublisher.PublishCharacters(filteredCharacters)
	if err != nil {
		return err
	}

	return nil
}

func deleteDuplicatedCharacters(log *logrus.Logger, characters []*models.Character) []*models.Character {
	filteredCharacters := make([]*models.Character, 0)

	for _, character := range characters {
		_, found := lo.Find(filteredCharacters, func(c *models.Character) bool {
			if c.Title == character.Title {
				log.Printf("duplicated character found: [%s]", character.Title)
				return true
			}
			return false
		})

		if !found {
			filteredCharacters = append(filteredCharacters, character)
		}
	}

	return filteredCharacters
}
