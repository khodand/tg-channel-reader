package fire

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

var ErrNotFound = errors.New("doc not found")

func NewClient(ctx context.Context, cred string) (*firestore.Client, error) {
	opts := option.WithCredentialsFile(cred)
	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create firebase app")
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create firestore client")
	}
	return client, nil
}
