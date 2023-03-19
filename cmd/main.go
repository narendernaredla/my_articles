package main

import (
	"my_blogs/cmd/server"
	"my_blogs/utils"
)

func main() {
	utils.Logger.Info("Starting server....")

	server := server.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		utils.Logger.Fatalf("Failed to start server: %v", err)
	}
}
