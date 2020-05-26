package category

import (
	"context"
	"io"

	"gocloud.dev/docstore"
)

func (c *category) list(ctx context.Context) ([]Category, error) {
	result := make([]Category, 0)
	iter := c.coll.Query().OrderBy("Name", docstore.Ascending).Get(ctx)
	defer iter.Stop()
	for {
		var p Category
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

func (c *category) add(ctx context.Context, data *Category) (CategoryID, error) {
	err := c.coll.Create(ctx, data)
	if err != nil {
		return "", err
	}
	return data.ID, nil
}
