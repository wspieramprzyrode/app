package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/wspieramprzyrode/mobile/inventory/datastore"
	"gocloud.dev/docstore"
)

type category struct {
	coll *docstore.Collection
}

func New(ctx context.Context) (*category, error) {
	categoryColl, err := datastore.New(ctx, "firebase", "wspieramprzyrode", "categories", "ID")
	if err != nil {
		return nil, err
	}
	return &category{coll: categoryColl}, nil
}

func (c *category) Close() {
	c.coll.Close()
}

type InventoryCategoryCreator interface {
	CreateInventoryCategory(ctx context.Context, data *Category) (CategoryID, error)
}

func (c *category) CreateInventoryCategory(ctx context.Context, data *Category) (CategoryID, error) {
	id := CategoryID(uuid.New().String())
	data.ID = id
	return c.add(ctx, data)
}
