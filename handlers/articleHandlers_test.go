//go:build !integration
// +build !integration

package handlers

import (
	"errors"
	"my_blogs/mocks"
	"my_blogs/models"
	"my_blogs/models/db"
	"my_blogs/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateArticle(t *testing.T) {
	articleModelMock := mocks.ArticleMock{}
	articleModelMock.On("Create", mock.Anything).Return("testid1", nil)

	errArticleModelMock := mocks.ArticleMock{}
	errArticleModelMock.On("Create", mock.Anything).Return(mock.Anything, errors.New("Failed"))

	tests := []struct {
		name           string
		payload        *strings.Reader
		expectedStatus int
		expectedResp   string
		articleModel   mocks.ArticleMock
	}{
		{
			name:           "Should returns 201 success",
			payload:        strings.NewReader(`{ "title":"adsads", "content":"adsads", "author":"adsads" }`),
			expectedStatus: http.StatusCreated,
			expectedResp:   "testid1",
			articleModel:   articleModelMock,
		},
		{
			name:           "Should returns 400 bad request",
			payload:        strings.NewReader(``),
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "",
			articleModel:   articleModelMock,
		},
		{
			name:           "Should returns 400 bad request with missing fields",
			payload:        strings.NewReader(`{ "content":"adsads", "author":"adsads" }`),
			expectedStatus: http.StatusBadRequest,
			expectedResp:   "",
			articleModel:   articleModelMock,
		},
		{
			name:           "Should returns 500 error",
			payload:        strings.NewReader(`{ "title":"adsads", "content":"adsads", "author":"adsads" }`),
			expectedStatus: http.StatusInternalServerError,
			expectedResp:   "",
			articleModel:   errArticleModelMock,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := db.NewArticleModel(test.articleModel)
			articleService := service.NewArticleService(articleModel)
			articlehandler := NewArticleHandler(articleService)
			handler := http.HandlerFunc(articlehandler.CreateArticle)
			response := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/articles", test.payload)
			if err != nil {
				t.Fatal(err)
			}
			handler.ServeHTTP(response, req)
			assert.Equal(t, test.expectedStatus, response.Code)
		})
	}
}

func TestGetArticlesById(t *testing.T) {
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
	articleModelMock := mocks.ArticleMock{}
	articleModelMock.On("FindOne", mock.Anything).Return(mockData, nil)

	noDocArticleModelMock := mocks.ArticleMock{}
	noDocArticleModelMock.On("FindOne", mock.Anything).Return(models.ArticleModel{}, nil)

	errarticleModelMock1 := mocks.ArticleMock{}
	errarticleModelMock1.On("FindOne", mock.Anything).Return(models.ArticleModel{}, errors.New("no document found"))

	tests := []struct {
		name           string
		articleId      string
		expectedStatus int
		articleModel   mocks.ArticleMock
	}{
		{
			name:           "Should returns 200 success",
			articleId:      `64155a8ead22c56000981fd0`,
			expectedStatus: http.StatusOK,
			articleModel:   articleModelMock,
		},
		{
			name:           "Should returns 200 no document found",
			articleId:      `64155a8ead22c56000981fd0`,
			expectedStatus: http.StatusOK,
			articleModel:   noDocArticleModelMock,
		},
		{
			name:           "Should returns 400 bad request",
			articleId:      ``,
			expectedStatus: http.StatusBadRequest,
			articleModel:   articleModelMock,
		},
		{
			name:           "Should returns 500",
			articleId:      `64155a8ead22c560ss00981fd0`,
			expectedStatus: http.StatusInternalServerError,
			articleModel:   errarticleModelMock1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := db.NewArticleModel(test.articleModel)
			articleService := service.NewArticleService(articleModel)
			articlehandler := NewArticleHandler(articleService)
			handler := http.HandlerFunc(articlehandler.GetArticlesById)
			response := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/articles", nil)
			req = mux.SetURLVars(req, map[string]string{"articleId": test.articleId})
			if err != nil {
				t.Fatal(err)
			}
			handler.ServeHTTP(response, req)
			assert.Equal(t, test.expectedStatus, response.Code)
		})
	}
}

func TestGetArticles(t *testing.T) {
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
	articleModelMock := mocks.ArticleMock{}
	articleModelMock.On("FindAll", mock.Anything).Return(mockData, nil)

	errarticleModelMock1 := mocks.ArticleMock{}
	errarticleModelMock1.On("FindAll", mock.Anything).Return([]models.ArticleModel{}, nil)

	errarticleModelMock2 := mocks.ArticleMock{}
	errarticleModelMock2.On("FindAll", mock.Anything).Return([]models.ArticleModel{}, errors.New("no documents found"))

	tests := []struct {
		name           string
		expectedStatus int
		articleModel   mocks.ArticleMock
	}{
		{
			name:           "Should returns 200 success",
			expectedStatus: http.StatusOK,
			articleModel:   articleModelMock,
		},
		{
			name:           "Should returns 200 no documents found",
			expectedStatus: http.StatusOK,
			articleModel:   errarticleModelMock1,
		},
		{
			name:           "Should returns 500",
			expectedStatus: http.StatusInternalServerError,
			articleModel:   errarticleModelMock2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleModel := db.NewArticleModel(test.articleModel)
			articleService := service.NewArticleService(articleModel)
			articlehandler := NewArticleHandler(articleService)
			handler := http.HandlerFunc(articlehandler.GetArticles)
			response := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/articles", nil)
			if err != nil {
				t.Fatal(err)
			}
			handler.ServeHTTP(response, req)
			assert.Equal(t, test.expectedStatus, response.Code)
		})
	}
}
