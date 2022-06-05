package usecase

type LOTRInformation interface {
	PublishCharactersInformation(filePath string) error
}
