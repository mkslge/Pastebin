// database/mongo.go
package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongo(ctx context.Context) (*mongo.Client, error) {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")

	if username == "" || password == "" {
		return nil, fmt.Errorf("MONGO_USERNAME and MONGO_PASSWORD must be set")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().
		ApplyURI("mongodb+srv://sharedcalendar.xmyrcke.mongodb.net/?appName=SharedCalendar").
		SetAuth(options.Credential{
			Username: username,
			Password: password,
		}).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
