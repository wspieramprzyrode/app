package objects

import (
	"context"
	"io"
)

func (c *objects) list(ctx context.Context) ([]InventoryObject, error) {
	result := make([]InventoryObject, 0)
	iter := c.coll.Query().Get(ctx)
	defer iter.Stop()
	for {
		var p InventoryObject
		err := iter.Next(ctx, &p)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		} else {
			result = append(result, p)
		}
	}
	return result, nil
}

func (c *objects) add(ctx context.Context, data *InventoryObject) (InventoryObjectID, error) {
	err := c.coll.Create(ctx, data)
	if err != nil {
		return "", err
	}
	return data.ID, nil
}
