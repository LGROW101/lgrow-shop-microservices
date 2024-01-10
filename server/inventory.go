package server

import (
	"github.com/LGROW101/LGROW-Microservices/modules/inventory/inventoryHandler"
	"github.com/LGROW101/LGROW-Microservices/modules/inventory/inventoryRepository"
	"github.com/LGROW101/LGROW-Microservices/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health Checr
	inventory.GET("/health", s.healthCheckService)
}
