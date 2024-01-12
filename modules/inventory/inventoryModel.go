package inventory

import (
	"github.com/LGROW101/LGROW-Microservices/modules/item"
	"github.com/LGROW101/LGROW-Microservices/modules/models"
)

type (
	UpdateInventoryReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}
	ItemInInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}
	PlayerInventory struct {
		PlayerId string `json:"player_id"`
		*models.PaginateRes
	}
)
