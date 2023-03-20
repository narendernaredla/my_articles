package db

import (
	"errors"
	"my_blogs/mocks"
	"my_blogs/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreate(t *testing.T) {
	collectionHelperMock := mocks.ArticleMock{}
	collectionHelperMock.On("Create", mock.Anything).Return("testid1", nil)

	errCollectionHelperMock := mocks.ArticleMock{}
	errCollectionHelperMock.On("Create", mock.Anything).Return(nil, errors.New("failed"))

	tests := []struct {
		name          string
		newArticle    *models.ArticleModel
		expectedId    string
		expectedError error
		model         mocks.ArticleMock
	}{
		{
			name: "Should return inserted id",
			newArticle: &models.ArticleModel{
				Title:   "test",
				Author:  "test",
				Content: "test content",
			},
			expectedId:    "testid1",
			expectedError: nil,
			model:         collectionHelperMock,
		},
		{
			name: "Should throw error",
			newArticle: &models.ArticleModel{
				Title:   "test",
				Author:  "test",
				Content: "test content",
			},
			expectedId:    "",
			expectedError: errors.New("failed to create article"),
			model:         errCollectionHelperMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := NewArticleModel(test.model)
			expectedID, err := articleModel.Create(test.newArticle)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedId, expectedID)
		})
	}
}

func TestGetById(t *testing.T) {
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

	noDocCollectionHelperMock := mocks.ArticleMock{}
	noDocCollectionHelperMock.On("FindOne", mock.Anything).Return(models.ArticleModel{}, nil)

	errCollectionHelperMock := mocks.ArticleMock{}
	errCollectionHelperMock.On("FindOne", mock.Anything).Return(nil, errors.New("failed"))

	tests := []struct {
		name          string
		newArticle    *models.ArticleModel
		id            string
		expectedError error
		model         mocks.ArticleMock
	}{
		{
			name:          "Should return article",
			id:            "64155a8ead22c56000981fd0",
			expectedError: nil,
			model:         collectionHelperMock,
		},
		{
			name:          "Should no documet",
			id:            "64155a8ead22c56000981fd0",
			expectedError: errors.New("no document found"),
			model:         noDocCollectionHelperMock,
		},
		{
			name:          "Should throw error",
			id:            "",
			expectedError: errors.New("failed to read article id"),
			model:         collectionHelperMock,
		},

		{
			name:          "Should throw error ee",
			id:            "64155a8ead22c56000981fd0",
			expectedError: errors.New("failed to read article"),
			model:         errCollectionHelperMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := NewArticleModel(test.model)
			_, err := articleModel.GetById(test.id)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestGetAll(t *testing.T) {
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

	noDocCollectionHelperMock := mocks.ArticleMock{}
	noDocCollectionHelperMock.On("FindAll", mock.Anything).Return([]models.ArticleModel{}, nil)

	errCollectionHelperMock := mocks.ArticleMock{}
	errCollectionHelperMock.On("FindAll", mock.Anything).Return(nil, errors.New("failed"))

	tests := []struct {
		name          string
		newArticle    *models.ArticleModel
		expectedError error
		model         mocks.ArticleMock
	}{
		{
			name:          "Should return []article",
			expectedError: nil,
			model:         collectionHelperMock,
		},
		{
			name:          "Should no documet",
			expectedError: errors.New("no documents found"),
			model:         noDocCollectionHelperMock,
		},
		{
			name:          "Should throw error ee",
			expectedError: errors.New("failed to read articles"),
			model:         errCollectionHelperMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := NewArticleModel(test.model)
			_, err := articleModel.GetAll()
			assert.Equal(t, test.expectedError, err)
		})
	}
}
