package models

import (
	"my_blogs/models/db"
	"my_blogs/utils"
)

type GetArticleResponseModel struct {
	CreateArticleResponseModel
	Title   string
	Content string
	Author  string
}

type GetAllArticleResponseModel struct {
	GetArticleResponseModel
}

func GetArticleById(id string) (GetArticleResponseModel, error) {
	utils.Logger.Info("GetArticleResponseModel::GetArticleById() :: Entered")
	getArticleResp := GetArticleResponseModel{}
	article := &db.Article{Id: id}
	getArticleByIdResp, _ := article.GetById()

	utils.Logger.Info("GetArticleResponseModel::Create() :: Returing respose to caller,", getArticleByIdResp)
	utils.Logger.Info("GetArticleResponseModel::Create() :: Returing respose to caller")
	return getArticleResp, nil
}

func GetAllArticles() ([]GetAllArticleResponseModel, error) {
	utils.Logger.Info("GetArticleResponseModel::GetArticleById() :: Entered")
	var getAllArticleObj []GetAllArticleResponseModel
	article := &db.Article{}
	getAllArticleResp, _ := article.GetAll()
	utils.Logger.Info("GetArticleResponseModel::Create() :: Returing respose to caller", getAllArticleResp)
	return getAllArticleObj, nil
}
