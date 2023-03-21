//go:build integration
// +build integration

package handlers

import (
	"my_blogs/mocks"
	"my_blogs/models/db"
	"my_blogs/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateArticleInt(t *testing.T) {
	tests := []struct {
		name           string
		payload        *strings.Reader
		expectedStatus int
		expectedResp   string
		articleModel   mocks.ArticleMock
	}{
		{
			name:           "Should returns 201 success",
			payload:        strings.NewReader(`{ "title":"test", "content":"test", "author":"test" }`),
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Should returns 400 Bad request missing fields",
			payload:        strings.NewReader(`{ "content":"test", "author":"test" }`),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Should returns 400 Bad request",
			payload:        strings.NewReader(``),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleCollection := db.NewCollectionHelper()
			articleModel := db.NewArticleModel(articleCollection)
			articleService := service.NewArticleService(articleModel)
			articleHandlers := NewArticleHandler(articleService)
			handler := http.HandlerFunc(articleHandlers.CreateArticle)
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

func TestGetArticlesByIdInt(t *testing.T) {

	tests := []struct {
		name           string
		articleId      string
		expectedStatus int
	}{
		{
			name:           "Should returns 200 no document found",
			articleId:      `64155a8ead22c56000981fd0`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Should returns 400 bad request",
			articleId:      ``,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleCollection := db.NewCollectionHelper()
			articleModel := db.NewArticleModel(articleCollection)
			articleService := service.NewArticleService(articleModel)
			articleHandlers := NewArticleHandler(articleService)
			handler := http.HandlerFunc(articleHandlers.GetArticlesById)
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

func TestGetArticlesInt(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Should returns 200 success",
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			articleCollection := db.NewCollectionHelper()
			articleModel := db.NewArticleModel(articleCollection)
			articleService := service.NewArticleService(articleModel)
			articleHandlers := NewArticleHandler(articleService)
			handler := http.HandlerFunc(articleHandlers.GetArticles)
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
