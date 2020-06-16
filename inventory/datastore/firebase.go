package datastore

import (
	"context"

	firestore "cloud.google.com/go/firestore/apiv1"
	"gocloud.dev/docstore/gcpfirestore"
	"gocloud.dev/gcp"
)

func newFirebaseClient(ctx context.Context) (*firestore.Client, error) {
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	client, _, err := gcpfirestore.Dial(ctx, creds.TokenSource)
	if err != nil {
		return nil, err
	}
	return client, nil

}
