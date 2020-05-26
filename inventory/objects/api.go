package objects

import (
	"context"

	"github.com/google/uuid"
	"github.com/wspieramprzyrode/mobile/inventory/datastore"
	"gocloud.dev/docstore"
)

type objects struct {
	coll *docstore.Collection
}

func New(ctx context.Context) (*objects, error) {
	coll, err := datastore.New(ctx, "firebase", "wspieramprzyrode", "inventory_objects", "ID")
	if err != nil {
		return nil, err
	}
	return &objects{coll: coll}, nil
}

func (o *objects) Close() {
	o.coll.Close()
}

type InventoryObjectCreator interface {
	CreateInventoryObject(ctx context.Context, data *InventoryObject) (InventoryObjectID, error)
}

func (o *objects) CreateInventoryObject(ctx context.Context, data *InventoryObject) (InventoryObjectID, error) {
	id := InventoryObjectID(uuid.New().String())
	data.ID = id
	return o.add(ctx, data)
}
