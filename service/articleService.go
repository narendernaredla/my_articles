package service

import (
	"my_blogs/models"
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
	Id string `json:"id,omitempty"`
}

type GetArticleResponseModel struct {
	CreateArticleResponseModel
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

type IArticleService interface {
	Create(articleReq *CreateArticleRequestModel) (CreateArticleResponseModel, error)
	Validate(articleReq *CreateArticleRequestModel) error
	GetArticleById(id string) (GetArticleResponseModel, error)
	GetAllArticles() ([]GetArticleResponseModel, error)
}

type ArticleService struct {
	articleModel db.IArticle
}

func NewArticleService(articleModel db.IArticle) IArticleService {
	return &ArticleService{
		articleModel: articleModel,
	}
}

func (articleSvc *ArticleService) Validate(articleReq *CreateArticleRequestModel) error {
	return validator.New().Struct(articleReq)
}

func (articleSvc *ArticleService) Create(articleReq *CreateArticleRequestModel) (CreateArticleResponseModel, error) {
	utils.Logger.Info("ArticleService::Create() :: Entered")
	article := &models.ArticleModel{}
	article.Title = articleReq.Title
	article.Author = articleReq.Author
	article.Content = articleReq.Content
	articleId, err := articleSvc.articleModel.Create(article)
	utils.Logger.Info("ArticleService::Create() :: Returing respose to caller")
	return CreateArticleResponseModel{Id: articleId}, err
}

func (articleSvc *ArticleService) GetArticleById(id string) (GetArticleResponseModel, error) {
	utils.Logger.Info("ArticleService::GetArticleById() :: Entered")
	getArticleByIdResp, err := articleSvc.articleModel.GetById(id)
	var getArticleResp = GetArticleResponseModel{}
	if err == nil {
		getArticleResp = ConvertToViewObject(getArticleByIdResp)
	}
	utils.Logger.Infof("ArticleService::GetArticleById() :: Returing respose to caller ArticleModel: %v", getArticleResp)
	return getArticleResp, err
}

func (articleSvc *ArticleService) GetAllArticles() ([]GetArticleResponseModel, error) {
	utils.Logger.Info("ArticleService::GetAllArticles() :: Entered")
	var getAllArticleObj = []GetArticleResponseModel{}
	getAllArticleResp, err := articleSvc.articleModel.GetAll()
	for _, val := range getAllArticleResp {
		getArticleResp := ConvertToViewObject(val)
		getAllArticleObj = append(getAllArticleObj, getArticleResp)
	}
	utils.Logger.Info("ArticleService::GetAllArticles() :: Returing respose to caller")
	return getAllArticleObj, err
}

func ConvertToViewObject(dbObject models.ArticleModel) GetArticleResponseModel {
	utils.Logger.Info("ArticleService::ConvertToViewObject() :: Entered")
	getArticleResp := GetArticleResponseModel{}
	getArticleResp.Id = dbObject.Id.Hex()
	getArticleResp.Author = dbObject.Author
	getArticleResp.Content = dbObject.Content
	getArticleResp.Title = dbObject.Title
	utils.Logger.Info("ArticleService::ConvertToViewObject() :: Returning caller")
	return getArticleResp
}
