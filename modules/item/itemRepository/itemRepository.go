package itemRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemRepositoryService interface{}
	itemRespository       struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepositoryService {
	return &itemRespository{db}
}
func (r *itemRespository) itemDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}
