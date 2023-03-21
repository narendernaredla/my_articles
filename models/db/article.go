package db

import (
	"context"
	"errors"
	"my_blogs/models"
	"my_blogs/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IArticle interface {
	Create(article *models.ArticleModel) (string, error)
	GetById(id string) (models.ArticleModel, error)
	GetAll() ([]models.ArticleModel, error)
}

type Article struct {
	articleCollection ICollectionHelper
}

func NewArticleModel(collection ICollectionHelper) IArticle {
	return &Article{
		articleCollection: collection,
	}
}

func (a *Article) Create(article *models.ArticleModel) (string, error) {
	utils.Logger.Info("articleModel::Create() :: Entered")
	insertId, err := a.articleCollection.Create(context.TODO(), article)
	if err != nil {
		utils.Logger.Errorf("articleModel::Create() :: Error while creating article. %v", err)
		return "", errors.New("failed to create article")
	}
	utils.Logger.Info("articleModel::Create() :: returning caller")
	return insertId, nil
}

func (a *Article) GetById(id string) (models.ArticleModel, error) {
	utils.Logger.Info("articleModel::GetById() :: Entered")
	var article models.ArticleModel
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.Logger.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
		return article, errors.New("failed to read article id")
	}
	article, err = a.articleCollection.FindOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		utils.Logger.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
		return article, errors.New("failed to read article")
	}
	if article.Title == "" {
		utils.Logger.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
		return article, errors.New("no document found")
	}
	utils.Logger.Info("articleModel::GetById() :: returning caller")
	return article, err
}

func (a *Article) GetAll() ([]models.ArticleModel, error) {
	utils.Logger.Info("articleModel::GetAll() :: Entered")
	var allArticles []models.ArticleModel
	allArticles, err := a.articleCollection.FindAll(context.TODO())
	if err != nil {
		utils.Logger.Errorf("articleModel::GetAll() :: Error while reading articles, %v", err)
		return allArticles, errors.New("failed to read articles")
	}
	if len(allArticles) <= 0 {
		err = errors.New("no documents found")
	}
	utils.Logger.Info("articleModel::GetAll() :: returning caller")
	return allArticles, err
}
