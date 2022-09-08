package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Client, err error) {
	log.Println("start create client")

	mongoDBURL := fmt.Sprintf("mongodb://%s:%s", host, port)

	clientOptions := options.Client().ApplyURI(mongoDBURL)

	if username != "" && password != "" {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			Username:   username,
			Password:   password,
			AuthSource: authDB,
		})
	}

	//connect
	log.Println("connect to db")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed connect to db error: %v", err)
	}

	//ping
	log.Println("ping to db")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed ping db error: %v", err)
	}

	return client, nil
}
