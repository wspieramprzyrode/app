package datastore

import (
	"context"
	"errors"

	"gocloud.dev/docstore"
	"gocloud.dev/docstore/gcpfirestore"
)

func New(ctx context.Context, storeType, projectName, collectionName, IdField string) (*docstore.Collection, error) {
	switch storeType {
	case "firebase":
		client, err := newFirebaseClient(ctx)
		if err != nil {
			return nil, err
		}
		resourceID := gcpfirestore.CollectionResourceID(projectName, collectionName)
		return gcpfirestore.OpenCollection(client, resourceID, IdField, nil)

	default:
		return nil, errors.New("Bad datastore type")
	}
}
