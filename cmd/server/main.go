package main

import (
	"my_blogs/handlers"
	"my_blogs/utils"
	"net/http"
	"os"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	utils.InitLogger()

	utils.Logger.Info("Started.....")

	articleHandlers := handlers.NewArticleHandler()

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

	err := http.ListenAndServe(":"+port, gorillaHandlers.CORS(headers, methods, origins)(router))
	if err != nil {
		utils.Logger.Fatalf("Failed to start server: %v", err)
	}

}
