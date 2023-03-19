package db

import (
	"context"
	"errors"
	"my_blogs/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Author  string             `json:"author,omitempty" bson:"author,omitempty"`
}

func (a *Article) Create() (string, error) {
	utils.Logger.Info("articleModel::Create() :: Entered")
	articlesCollection := GetDB().Collection("articles")
	result, err := articlesCollection.InsertOne(context.TODO(), a)
	if err != nil {
		utils.Logger.Errorf("articleModel::Create() :: Error while creating article. %v", err)
		return "", errors.New("failed to create article")
	}
	utils.Logger.Info("articleModel::Create() :: returning caller")
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *Article) GetById(id string) (Article, error) {
	utils.Logger.Info("articleModel::GetById() :: Entered")
	var article Article
	articlesCollection := GetDB().Collection("articles")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.Logger.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
		return article, errors.New("failed to read article id")
	}
	err = articlesCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&article)
	if err != nil {
		utils.Logger.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
		return article, errors.New("no document found")
	}
	utils.Logger.Info("articleModel::GetById() :: returning caller")
	return article, err
}

func (a *Article) GetAll() ([]Article, error) {
	utils.Logger.Info("articleModel::GetAll() :: Entered")
	var allArticles []Article
	articlesCollection := GetDB().Collection("articles")
	cur, err := articlesCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		utils.Logger.Errorf("articleModel::GetAll() :: Error while reading articles, %v", err)
		return allArticles, errors.New("failed to read articles")
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var article Article
		err := cur.Decode(&article)
		if err != nil {
			utils.Logger.Errorf("articleModel::GetAll() :: Error while processing articles, %v", err)
			return allArticles, errors.New("failed to read articles")
		}

		allArticles = append(allArticles, article)
	}
	utils.Logger.Info("articleModel::GetAll() :: returning caller")
	return allArticles, errors.New("Failed")
}
