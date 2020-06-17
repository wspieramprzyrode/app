package objects

import "github.com/wspieramprzyrode/mobile/inventory/category"

type InventoryObjectID string

type InventoryObject struct {
	ID               InventoryObjectID   `json:"id,omitempty"`
	CategoryID       category.CategoryID `json:"category_id" binding:"required"`
	Coordinates      Coordinates         `json:"coordinates" binding:"required"`
	CreatedBy        string              `json:"-"`
	DocstoreRevision interface{}         `json:"-"`
}

type Coordinates struct {
	Lat float32 `json:"lat" binding:"required"`
	Lng float32 `json:"lng" binding:"required"`
}
