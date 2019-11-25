package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const keyjson = "pkg/conf/key.json"

func OpenFirebase() *firestore.Client {
	opt := option.WithCredentialsFile(keyjson)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Errorf("error initializing client: %v", err)
	}
	return client
}
