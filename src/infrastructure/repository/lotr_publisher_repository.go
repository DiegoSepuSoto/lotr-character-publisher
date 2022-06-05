package repository

import "DiegoSepuSoto/lotr-character-publisher/src/domain/models"

type LOTRPublisherRepository interface {
	PublishCharacters(character []*models.Character) error
}
