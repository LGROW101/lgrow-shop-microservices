package playerRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PlayerRespositoryService interface{}

	playerRespository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRespositoryService {
	return &playerRespository{db: db}
}

func (r *playerRespository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}
