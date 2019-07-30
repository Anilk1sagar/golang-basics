package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

func InitFirebase() *db.Client {

	/* ========================== Initialize Firebase Admin ======================= */

	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://gobaseapp-82fd5.firebaseio.com",
	}
	opt := option.WithCredentialsFile("gobaseapp-82fd5-firebase-adminsdk.json")

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal("Error initializing Firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal("Error initializing Firebase database: ", err)
	}

	// fmt.Println("** Firebase Initialize: ** ", app, client)

	return client
}
