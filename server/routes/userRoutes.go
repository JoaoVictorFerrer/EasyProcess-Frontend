package routes

import (
	"EasyProcess/handlers"
	"EasyProcess/storage"
	"EasyProcess/utils"
	"net/http"
)

func SetupUserRoutes(router *http.ServeMux, storage storage.StorageInterface) {
	// GET
	router.HandleFunc("GET /api/v1/user/me", utils.ProtectedHTTPHandleFunc(handlers.HandleGetUserMe, storage))
}
