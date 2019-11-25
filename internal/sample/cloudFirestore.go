package main

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// クライアント接続
	opt := option.WithCredentialsFile("key.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Errorf("error initializing client: %v", err)
	}
	defer client.Close()
	fmt.Println("Connection done")

	// 値の取得
	collection := client.Collection("コレクションID")
	doc := collection.Doc("ドキュメントID")
	field, err := doc.Get(ctx)
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	data := field.Data()
	for key, value := range data {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
