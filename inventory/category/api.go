package category

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wspieramprzyrode/mobile/inventory/datastore"
	"gocloud.dev/docstore"
)

type category struct {
	coll *docstore.Collection
}

func New(ctx context.Context, projectID, appEnv string) (*category, error) {
	categoryColl, err := datastore.New(ctx, "firebase", projectID, fmt.Sprintf("inventory_categories_%s", appEnv), "ID")
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
