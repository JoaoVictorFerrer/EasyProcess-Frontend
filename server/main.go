package main

import (
	"EasyProcess/routes"
	"EasyProcess/storage"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load local config file in DEV environment as
	if os.Getenv("ENV") != "PROD" {
		godotenv.Load("config.env")
	}
	mongoClient := ConnectToDatabase()
	storage := storage.InitStorage(mongoClient)
	router := http.NewServeMux()
	routes.SetupAuthRoutes(router, storage)
	routes.SetupUserRoutes(router, storage)
	StartServer(router)
}
