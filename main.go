package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	_ = godotenv.Load()

	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoUsername := os.Getenv("MONGO_USERNAME")
	if mongoUsername == "" || mongoPassword == "" {
		log.Fatal("MONGO_USERNAME and MONGO_PASSWORD must be set in your environment or .env file")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	uri := "mongodb+srv://sharedcalendar.xmyrcke.mongodb.net/?appName=SharedCalendar"
	opts := options.Client().
		ApplyURI(uri).
		SetAuth(options.Credential{
			Username: mongoUsername,
			Password: mongoPassword,
		}).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
