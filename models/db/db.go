package db

import (
	"context"
	"log"
	"my_blogs/utils"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB Client Instance
var db *mongo.Client
var once sync.Once

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/my_articles?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	utils.Logger.Info("Connected to DB")
	return client
}

func GetDB() *mongo.Database {
	once.Do(func() {
		db = ConnectDB()
	})
	return db.Database("my_articles")
}
