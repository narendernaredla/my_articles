package handlers

import (
	"encoding/json"
	"my_blogs/service"
	"my_blogs/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type ArticleHandler struct {
	articleService service.IArticleService
}

type CreateArticleRequestModel struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}

func NewArticleHandler(articleService service.IArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

func (a *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("articleHandlers::CreateArticle() :: Entered", a.articleService)
	createArticleReq := &service.CreateArticleRequestModel{}
	err := json.NewDecoder(r.Body).Decode(createArticleReq)
	if err != nil {
		utils.Logger.Errorf("articleHandlers::CreateArticle() :: Failed to read payload. %v", err)
		resp := utils.FormatResponse(http.StatusBadRequest, "Failed to read data", nil)
		utils.SendResponse(w, resp, http.StatusBadRequest)
		return
	}
	utils.Logger.Infof("articleHandlers::CreateArticle() :: Received Payload: %+v", createArticleReq)

	if err := a.articleService.Validate(createArticleReq); err != nil {
		utils.Logger.Errorf("articleHandlers::CreateArticle() :: Error validating request. %v", err)
		resp := utils.FormatResponse(http.StatusBadRequest, "Missing required fields", nil)
		utils.SendResponse(w, resp, http.StatusBadRequest)
		return
	}

	createArticleResp, err := a.articleService.Create(createArticleReq)

	if err != nil {
		utils.Logger.Errorf("articleHandlers::CreateArticle() :: Error while creating article: %v", err)
		resp := utils.FormatResponse(http.StatusInternalServerError, err.Error(), nil)
		utils.SendResponse(w, resp, http.StatusInternalServerError)
		return
	}

	utils.Logger.Infof("articleHandlers::CreateArticle() :: Sending response to user %v", createArticleResp)
	resp := utils.FormatResponse(http.StatusCreated, "Success", createArticleResp)
	utils.SendResponse(w, resp, http.StatusCreated)
}

func (a *ArticleHandler) GetArticlesById(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("articleHandlers::GetArticlesById() :: Entered")
	params := mux.Vars(r)
	articleId := params["articleId"]
	utils.Logger.Infof("articleHandlers::GetArticlesById() :: Received articleId: %v", articleId)

	if len(articleId) <= 0 {
		utils.Logger.Errorf("articleHandlers::GetArticlesById() :: Missing articleId")
		resp := utils.FormatResponse(http.StatusBadRequest, "Missing articleId", nil)
		utils.SendResponse(w, resp, http.StatusBadRequest)
		return
	}
	getArticleResp, err := a.articleService.GetArticleById(articleId)
	if err != nil {
		utils.Logger.Errorf("articleHandlers::GetArticlesById() :: Error while reading article by id: %v, Error: %v", articleId, err)
		if err.Error() == "no document found" {
			resp := utils.FormatResponse(http.StatusOK, err.Error(), getArticleResp)
			utils.SendResponse(w, resp, http.StatusOK)
		} else {
			resp := utils.FormatResponse(http.StatusInternalServerError, err.Error(), getArticleResp)
			utils.SendResponse(w, resp, http.StatusInternalServerError)
		}
		return
	}
	utils.Logger.Info("articleHandlers::GetArticlesById() :: Sending response to user")
	resp := utils.FormatResponse(http.StatusOK, "Success", getArticleResp)
	utils.SendResponse(w, resp, http.StatusOK)
}

func (a *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("articleHandlers::GetArticles() :: Entered")
	getAllArticleResp, err := a.articleService.GetAllArticles()
	if err != nil {
		utils.Logger.Errorf("articleHandlers::GetArticles() :: Error while reading all articles: %v", err)
		if err.Error() == "no documents found" {
			resp := utils.FormatResponse(http.StatusOK, err.Error(), getAllArticleResp)
			utils.SendResponse(w, resp, http.StatusOK)
		} else {
			resp := utils.FormatResponse(http.StatusInternalServerError, err.Error(), getAllArticleResp)
			utils.SendResponse(w, resp, http.StatusInternalServerError)
		}
		return
	}
	utils.Logger.Info("articleHandlers::GetArticles() :: Sending response to user")
	resp := utils.FormatResponse(http.StatusOK, "Success", getAllArticleResp)
	utils.SendResponse(w, resp, http.StatusOK)
}
