//go:build integration
// +build integration

package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"my_blogs/models"
	"testing"
)

var insertedId string

func TestGetDB(t *testing.T) {
	client := GetDB()
	assert.NotNil(t, client)
}

func TestConnectDB(t *testing.T) {
	client := ConnectDB()
	assert.NotNil(t, client)
}

func TestCreate(t *testing.T) {
	helper := NewCollectionHelper()
	newArticle := &models.ArticleModel{
		Title:   "test",
		Author:  "test",
		Content: "test content",
	}
	id, err := helper.Create(context.TODO(), newArticle)
	insertedId = id
	assert.NoError(t, err)
	assert.NotEqual(t, "", id)
}

func TestFindOne(t *testing.T) {
	helper := NewCollectionHelper()
	objectID, err := primitive.ObjectIDFromHex(insertedId)
	if err != nil {
		t.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
	}
	article, err := helper.FindOne(context.TODO(), bson.M{"_id": objectID})
	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.Equal(t, "test", article.Title)

	article, err2 := helper.FindOne(context.TODO(), "test")
	assert.NotNil(t, err2)
}

func TestFindAll(t *testing.T) {
	helper := NewCollectionHelper()
	articles, err := helper.FindAll(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, articles)
	assert.Greater(t, len(articles), 0)
}
