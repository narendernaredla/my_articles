package handlers

import (
	"encoding/json"
	"my_blogs/models"
	"my_blogs/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type articleHandler struct{}

func NewArticleHandler() *articleHandler {
	return &articleHandler{}
}

func (a *articleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("articleHandlers::CreateArticle() :: Entered")
	createArticleReq := &models.CreateArticleRequestModel{}
	err := json.NewDecoder(r.Body).Decode(createArticleReq)
	if err != nil {
		utils.Logger.Error("articleHandlers::CreateArticle() :: Failed to read payload")
		resp := utils.FormatResponse(http.StatusBadRequest, "Failed to read data", nil)
		utils.SendResponse(w, resp, http.StatusBadRequest)
		return
	}
	utils.Logger.Infof("articleHandlers::CreateArticle() :: Received Payload: %+v", createArticleReq)

	if err := createArticleReq.Validate(); err != nil {
		utils.Logger.Error("articleHandlers::CreateArticle() :: Error validating request")
		resp := utils.FormatResponse(http.StatusBadRequest, "Missing required fields", nil)
		utils.SendResponse(w, resp, http.StatusBadRequest)
		return
	}

	createArticleResp, err := createArticleReq.Create()
	if err != nil {
		utils.Logger.Errorf("articleHandlers::CreateArticle() :: Error while creating article: %v", err)
		resp := utils.FormatResponse(http.StatusInternalServerError, err.Error(), nil)
		utils.SendResponse(w, resp, http.StatusInternalServerError)
		return
	}

	utils.Logger.Info("articleHandlers::CreateArticle() :: Sending response to user")
	resp := utils.FormatResponse(http.StatusCreated, "Success", createArticleResp)
	utils.SendResponse(w, resp, http.StatusCreated)
}

func (a *articleHandler) GetArticlesById(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("articleHandlers::GetArticlesById() :: Entered GetArticlesById()")
	params := mux.Vars(r)
	articleId := params["articleId"]
	utils.Logger.Infof("articleHandlers::GetArticlesById() :: Received articleId: %v", articleId)

	if len(articleId) <= 0 {
		utils.Logger.Error("articleHandlers::GetArticlesById() :: Missing articleId")
		resp := utils.FormatResponse(http.StatusBadRequest, "Missing articleId", nil)
		utils.SendResponse(w, resp, http.StatusBadRequest)
		return
	}
	getArticleResp, err := models.GetArticleById(articleId)
	if err != nil {
		utils.Logger.Errorf("articleHandlers::GetArticlesById() :: Error while reading article by id: %v, Error: %v", articleId, err)
		resp := utils.FormatResponse(http.StatusInternalServerError, err.Error(), nil)
		utils.SendResponse(w, resp, http.StatusInternalServerError)
		return
	}

	resp := utils.FormatResponse(http.StatusOK, "Success", getArticleResp)
	utils.SendResponse(w, resp, http.StatusOK)
}

func (a *articleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("articleHandlers::GetArticles() :: Entered GetArticles()")

	getAllArticleResp, err := models.GetAllArticles()
	if err != nil {
		utils.Logger.Errorf("articleHandlers::GetArticles() :: Error while reading all articles: %v", err)
		resp := utils.FormatResponse(http.StatusInternalServerError, err.Error(), nil)
		utils.SendResponse(w, resp, http.StatusInternalServerError)
		return
	}
	resp := utils.FormatResponse(http.StatusOK, "Success", getAllArticleResp)
	utils.SendResponse(w, resp, http.StatusOK)
}
