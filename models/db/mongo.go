package db

import (
	"context"
	"errors"
	"log"
	"my_blogs/models"
	"my_blogs/utils"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB Client Instance
var db *mongo.Client
var once sync.Once

type ICollectionHelper interface {
	Create(context.Context, *models.ArticleModel) (string, error)
	FindOne(context.Context, interface{}) (models.ArticleModel, error)
	FindAll(context.Context) ([]models.ArticleModel, error)
}

type CollectionHelper struct {
}

func ConnectDB() *mongo.Client {
	mongoURI := os.Getenv("MONGOURI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:rootpassword@localhost:27017/myarticles?retryWrites=true&w=majority"
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
	return db.Database("myarticles")
}

func NewCollectionHelper() ICollectionHelper {
	return &CollectionHelper{}
}

func (helper *CollectionHelper) Create(ctx context.Context, article *models.ArticleModel) (string, error) {
	result, err := GetDB().Collection("articles").InsertOne(ctx, article)
	if err != nil {
		utils.Logger.Errorf("CollectionHelper::Create :: Error while creating article %v", err)
		return "", errors.New("failed to create article")
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (helper *CollectionHelper) FindOne(ctx context.Context, filter interface{}) (models.ArticleModel, error) {
	var article models.ArticleModel
	err := GetDB().Collection("articles").FindOne(context.TODO(), filter).Decode(&article)
	if err != nil && err.Error() != "mongo: no documents in result" {
		utils.Logger.Errorf("CollectionHelper::FindOne :: Error while reading article %v", err)
		return article, err
	}
	return article, nil
}

func (helper *CollectionHelper) FindAll(ctx context.Context) ([]models.ArticleModel, error) {
	var allArticles []models.ArticleModel
	cur, err := GetDB().Collection("articles").Find(ctx, bson.M{})
	if err != nil {
		utils.Logger.Errorf("CollectionHelper::FindAll :: Error while reading articles %v", err)
		return allArticles, errors.New("failed to read articles")
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var article models.ArticleModel
		err := cur.Decode(&article)
		if err != nil {
			return allArticles, errors.New("failed to read articles")
		}

		allArticles = append(allArticles, article)
	}
	return allArticles, nil
}
