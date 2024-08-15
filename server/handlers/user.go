package handlers

import (
	"EasyProcess/serverErrors"
	"EasyProcess/storage"
	"EasyProcess/utils"
	"net/http"
)

func HandleGetUserMe(w http.ResponseWriter, r *http.Request, db storage.StorageInterface) {
	// Get token from header
	token := r.Header.Get("Bearer-Token")
	userEmail, err := utils.ExtractEmailFromToken(token)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to validate input data", serverErrors.GenericError)
		return
	}
	// Get user from db
	user, err := db.GetUserByEmail(userEmail)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to validate input data", serverErrors.GenericError)
		return
	}
	if user == nil {
		utils.SendError(w, http.StatusNotFound, "User does not exist", serverErrors.GenericError)
		return
	}
	// Send response
	utils.SendJSON(w, http.StatusOK, user)
}
