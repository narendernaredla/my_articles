package server

import (
	"my_blogs/handlers"
	"my_blogs/models/db"
	"my_blogs/service"
	"net/http"
	"os"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewServer() *http.Server {
	router := mux.NewRouter()

	articleCollection := db.NewCollectionHelper()
	articleModel := db.NewArticleModel(articleCollection)

	articleService := service.NewArticleService(articleModel)
	articleHandlers := handlers.NewArticleHandler(articleService)

	router.HandleFunc("/articles", articleHandlers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{articleId}", articleHandlers.GetArticlesById).Methods("GET")
	router.HandleFunc("/articles", articleHandlers.GetArticles).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	headers := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := gorillaHandlers.AllowedMethods([]string{"GET", "POST"})
	origins := gorillaHandlers.AllowedOrigins([]string{"*"})

	server := http.Server{
		Addr:    ":" + port,
		Handler: gorillaHandlers.CORS(headers, methods, origins)(router),
	}
	return &server
}
