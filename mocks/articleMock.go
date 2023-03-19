package mocks

import (
	"context"
	"my_blogs/models"

	"github.com/stretchr/testify/mock"
)

type ArticleMock struct {
	mock.Mock
}

func (_m ArticleMock) Create(ctx context.Context, article *models.ArticleModel) (string, error) {
	ret := _m.Called(article)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *models.ArticleModel) string); ok {
		r0 = rf(ctx, article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.ArticleModel) error); ok {
		r1 = rf(ctx, article)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (_m ArticleMock) FindOne(ctx context.Context, id interface{}) (models.ArticleModel, error) {
	ret := _m.Called(id)

	var r0 models.ArticleModel
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) models.ArticleModel); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.ArticleModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (_m ArticleMock) FindAll(ctx context.Context) ([]models.ArticleModel, error) {
	ret := _m.Called()

	var r0 []models.ArticleModel
	if rf, ok := ret.Get(0).(func(context.Context) []models.ArticleModel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ArticleModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}
