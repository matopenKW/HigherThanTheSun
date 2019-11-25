package main

import (
	"HigherThanTheSun/pkg/db"
	"context"
	"fmt"
)

func main() {
	firebaseExe()
}

func firebaseExe() {
	client := db.OpenFirebase()

	// 値の取得
	collection := client.Collection("users")
	doc := collection.Doc("BxX1IX8lz4OFaETOvRHM")
	field, err := doc.Get(context.Background())
	if err != nil {
		fmt.Errorf("error get data: %v", err)
	}
	data := field.Data()
	for key, value := range data {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
