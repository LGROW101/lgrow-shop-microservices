package authHandler

import (
	authPb "github.com/LGROW101/LGROW-Microservices/modules/auth/authPb"
	"github.com/LGROW101/LGROW-Microservices/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAutGrpchHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler {
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}

}
