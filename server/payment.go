package server

import (
	"github.com/LGROW101/LGROW-Microservices/modules/payment/paymentHandler"
	"github.com/LGROW101/LGROW-Microservices/modules/payment/paymentRepository"
	"github.com/LGROW101/LGROW-Microservices/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymentHandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	payment := s.app.Group("/payment_v1")

	// Health Checr
	payment.GET("", s.healthCheckService)
}
