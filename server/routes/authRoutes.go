package routes

import (
	"EasyProcess/handlers"
	"EasyProcess/storage"
	"EasyProcess/utils"
	"net/http"
)

func SetupAuthRoutes(router *http.ServeMux, storage storage.StorageInterface) {
	// POST
	router.HandleFunc("POST /api/v1/auth/login", utils.HTTPHandleFunc(handlers.HandleLogin, storage))
	router.HandleFunc("POST /api/v1/auth/signup", utils.HTTPHandleFunc(handlers.HandleSignup, storage))
	router.HandleFunc("POST /api/v1/auth/refresh", utils.HTTPHandleFunc(handlers.HandleRefreshToken, storage))
}
