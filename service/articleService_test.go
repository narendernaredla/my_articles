package service

import (
	"errors"
	"my_blogs/mocks"
	"my_blogs/models"
	"my_blogs/models/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidate(t *testing.T) {
	articleModel := &db.Article{}
	articleService := NewArticleService(articleModel)

	newArticle := &CreateArticleRequestModel{
		Title:   "test title",
		Author:  "test author",
		Content: "test content",
	}
	err := articleService.Validate(newArticle)
	assert.NoError(t, err)

	newArticle = &CreateArticleRequestModel{
		Title:  "test title",
		Author: "test author",
	}
	err = articleService.Validate(newArticle)
	assert.NotNil(t, err)
}

func TestCreateArticle(t *testing.T) {
	collectionHelperMock := mocks.ArticleMock{}
	collectionHelperMock.On("Create", mock.Anything).Return("testid1", nil)

	errCollectionHelperMock := mocks.ArticleMock{}
	errCollectionHelperMock.On("Create", mock.Anything).Return(nil, errors.New("failed"))

	tests := []struct {
		name          string
		newArticle    *CreateArticleRequestModel
		expectedId    string
		expectedError error
		model         mocks.ArticleMock
	}{
		{
			name: "Should return inserted id",
			newArticle: &CreateArticleRequestModel{
				Title:   "test",
				Author:  "test",
				Content: "test content",
			},
			expectedId:    "testid1",
			expectedError: nil,
			model:         collectionHelperMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := db.NewArticleModel(test.model)
			articleService := NewArticleService(articleModel)
			expectedID, err := articleService.Create(test.newArticle)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedId, expectedID.Id)
		})
	}
}

func TestGetArticleById(t *testing.T) {
	objectID, err := primitive.ObjectIDFromHex(`64155a8ead22c56000981fd0`)
	if err != nil {
		t.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
	}
	mockData := models.ArticleModel{
		Id:      objectID,
		Title:   `test title`,
		Author:  `test author`,
		Content: `test content`,
	}

	collectionHelperMock := mocks.ArticleMock{}
	collectionHelperMock.On("FindOne", mock.Anything).Return(mockData, nil)

	errCollectionHelperMock := mocks.ArticleMock{}
	errCollectionHelperMock.On("FindOne", mock.Anything).Return(nil, errors.New("failed"))

	tests := []struct {
		name          string
		id            string
		expectedId    string
		expectedError error
		model         mocks.ArticleMock
	}{
		{
			name:          "Should return article",
			id:            "64155a8ead22c56000981fd0",
			expectedId:    "64155a8ead22c56000981fd0",
			expectedError: nil,
			model:         collectionHelperMock,
		},
		{
			name:          "Should return error",
			id:            "64155a8ead22c56000981fd2",
			expectedId:    "",
			expectedError: errors.New("failed to read article"),
			model:         errCollectionHelperMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := db.NewArticleModel(test.model)
			articleService := NewArticleService(articleModel)
			expectedID, err := articleService.GetArticleById(test.id)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedId, expectedID.Id)
		})
	}
}

func TestGetAllArticles(t *testing.T) {
	objectID, err := primitive.ObjectIDFromHex(`64155a8ead22c56000981fd0`)
	if err != nil {
		t.Errorf("articleModel::GetById() :: Error while reading article. %v", err)
	}
	mockData := []models.ArticleModel{
		{
			Id:      objectID,
			Title:   `test title`,
			Author:  `test author`,
			Content: `test content`,
		},
	}

	collectionHelperMock := mocks.ArticleMock{}
	collectionHelperMock.On("FindAll", mock.Anything).Return(mockData, nil)

	errCollectionHelperMock := mocks.ArticleMock{}
	errCollectionHelperMock.On("FindAll", mock.Anything).Return(nil, errors.New("failed"))

	tests := []struct {
		name          string
		expectedError error
		model         mocks.ArticleMock
	}{
		{
			name:          "Should return article",
			expectedError: nil,
			model:         collectionHelperMock,
		},
		{
			name:          "Should return error",
			expectedError: errors.New("failed to read articles"),
			model:         errCollectionHelperMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := db.NewArticleModel(test.model)
			articleService := NewArticleService(articleModel)
			_, err := articleService.GetAllArticles()
			assert.Equal(t, test.expectedError, err)
		})
	}
}
