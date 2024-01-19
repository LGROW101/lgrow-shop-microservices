package server

import (
	"log"

	"github.com/LGROW101/LGROW-Microservices/modules/player/playerHandler"
	playerPb "github.com/LGROW101/LGROW-Microservices/modules/player/playerPb"
	"github.com/LGROW101/LGROW-Microservices/modules/player/playerRepository"
	"github.com/LGROW101/LGROW-Microservices/modules/player/playerUsecase"
	"github.com/LGROW101/LGROW-Microservices/pkg/grpccon"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUsl)
		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUsl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health Check
	player.GET("", s.healthCheckService)

	player.POST("/player/register", httpHandler.CreatePlayer)
	player.POST("/player/add-money", httpHandler.AddPlayerMoney)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	player.GET("/player/account/:player_id", httpHandler.GetPlayerSavingAccount)

}
