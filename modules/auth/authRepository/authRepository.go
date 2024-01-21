package authRepository

import (
	"context"
	"errors"
	"log"
	"time"

	playerPb "github.com/LGROW101/LGROW-Microservices/modules/player/playerPb"
	"github.com/LGROW101/LGROW-Microservices/pkg/grpccon"
	"github.com/LGROW101/LGROW-Microservices/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface{}
	authRepository        struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}
func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
func (r *authRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Player().CredentialSearch(ctx, req)
	if err != nil {
		log.Printf("Error: CredentialSearch failed: %s", err.Error())
		return nil, errors.New("error: email or password is incorrect")
	}

	return result, nil
}
