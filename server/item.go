package server

import (
	"github.com/LGROW101/LGROW-Microservices/modules/item/itemHandler"
	"github.com/LGROW101/LGROW-Microservices/modules/item/itemRepository"
	"github.com/LGROW101/LGROW-Microservices/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health Checr
	item.GET("/health", s.healthCheckService)
}
