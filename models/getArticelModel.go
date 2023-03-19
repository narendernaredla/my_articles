package models

import (
	"my_blogs/models/db"
	"my_blogs/utils"
)

type GetArticleResponseModel struct {
	CreateArticleResponseModel
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func GetArticleById(id string) (GetArticleResponseModel, error) {
	utils.Logger.Info("GetArticleModel::GetArticleById() :: Entered")
	article := &db.Article{}
	getArticleByIdResp, err := article.GetById(id)
	getArticleResp := ConvertToViewObject(getArticleByIdResp)
	utils.Logger.Infof("GetArticleModel::GetArticleById() :: Returing respose to caller Article: %v, Error: %v", getArticleByIdResp)
	return getArticleResp, err
}

func GetAllArticles() ([]GetArticleResponseModel, error) {
	utils.Logger.Info("GetArticleModel::GetAllArticles() :: Entered")
	var getAllArticleObj []GetArticleResponseModel
	article := &db.Article{}
	getAllArticleResp, _ := article.GetAll()
	for _, val := range getAllArticleResp {
		getArticleResp := ConvertToViewObject(val)
		getAllArticleObj = append(getAllArticleObj, getArticleResp)
	}
	utils.Logger.Info("GetArticleModel::GetAllArticles() :: Returing respose to caller", getAllArticleObj)
	return getAllArticleObj, nil
}

func ConvertToViewObject(dbObject db.Article) GetArticleResponseModel {
	utils.Logger.Info("GetArticleModel::ConvertToViewObject() :: Entered")
	getArticleResp := GetArticleResponseModel{}
	getArticleResp.Id = dbObject.Id.Hex()
	getArticleResp.Author = dbObject.Author
	getArticleResp.Content = dbObject.Content
	getArticleResp.Title = dbObject.Title
	utils.Logger.Info("GetArticleModel::ConvertToViewObject() :: Returning caller")
	return getArticleResp
}
