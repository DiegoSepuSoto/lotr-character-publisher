package lotr

import (
	"DiegoSepuSoto/lotr-character-publisher/src/domain/models"
	"DiegoSepuSoto/lotr-character-publisher/src/utils/logger"
	"os"
)

func (r *lotrJSONRepository) GetCharactersInformation(filePath string) ([]*models.Character, error) {
	log := logger.GetLogger("repository", "GetCharactersInformation")

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Errorf("can't read file [%s]", err.Error())
		return nil, err
	}
	log.Println("file read successfully")

	var charactersFile *models.CharactersFile

	err = json.Unmarshal(file, &charactersFile)
	if err != nil {
		log.Errorf("can't unmarshal file into variable [%s]", err.Error())
		return nil, err
	}
	log.Println("file unmarshalled successfully")

	return charactersFile.Characters, nil
}
