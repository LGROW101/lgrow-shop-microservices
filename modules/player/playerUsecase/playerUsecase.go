package playerUsecase

import "github.com/LGROW101/LGROW-Microservices/modules/player/playerRepository"

type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRespository playerRepository.PlayerRespositoryService
	}
)

func NewPlayerUsecase(playerRespository playerRepository.PlayerRespositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRespository: playerRespository}
}
