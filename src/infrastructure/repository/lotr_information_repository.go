package repository

import "DiegoSepuSoto/lotr-character-publisher/src/domain/models"

type LOTRInformationRepository interface {
	GetCharactersInformation(filePath string) ([]*models.Character, error)
}
