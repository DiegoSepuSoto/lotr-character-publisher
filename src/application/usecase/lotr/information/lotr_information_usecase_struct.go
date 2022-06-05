package information

import "DiegoSepuSoto/lotr-character-publisher/src/infrastructure/repository"

type lotrInformationUseCase struct {
	lotrInformation repository.LOTRInformationRepository
	lotrPublisher   repository.LOTRPublisherRepository
}

func NewLOTRInformationUseCase(lotrInformation repository.LOTRInformationRepository,
	lotrPublisher repository.LOTRPublisherRepository) *lotrInformationUseCase {
	return &lotrInformationUseCase{
		lotrInformation: lotrInformation,
		lotrPublisher:   lotrPublisher,
	}
}
