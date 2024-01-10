package server

import (
	"github.com/LGROW101/LGROW-Microservices/modules/auth/authHandler"
	"github.com/LGROW101/LGROW-Microservices/modules/auth/authRepository"
	"github.com/LGROW101/LGROW-Microservices/modules/auth/authUsecase"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAutGrpchHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health Check
	auth.GET("/health", s.healthCheckService)
}
