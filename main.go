package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, _, err = client.Collection("test").Add(ctx, map[string]interface{}{
		"name":    "jiro",
		"age":     30,
		"country": "Japan",
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

}
