package models

import (
	"my_blogs/models/db"
	"my_blogs/utils"

	"gopkg.in/go-playground/validator.v9"
)

type CreateArticleRequestModel struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}

type CreateArticleResponseModel struct {
	Id string
}

func (reqModel *CreateArticleRequestModel) Validate() error {
	return validator.New().Struct(reqModel)
}

func (reqModel *CreateArticleRequestModel) Create() (CreateArticleResponseModel, error) {
	utils.Logger.Info("createArticleRequestModel::Create() :: Entered")
	article := &db.Article{}
	article.Title = reqModel.Title
	article.Author = reqModel.Author
	article.Content = reqModel.Content
	articleId, err := article.Create()
	utils.Logger.Info("createArticleRequestModel::Create() :: Returing respose to caller")
	return CreateArticleResponseModel{Id: articleId}, err
}
