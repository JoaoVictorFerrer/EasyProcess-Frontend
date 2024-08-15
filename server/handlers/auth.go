package handlers

import (
	"EasyProcess/models"
	"EasyProcess/serverErrors"
	"EasyProcess/storage"
	"EasyProcess/utils"
	"encoding/json"
	"log"
	"net/http"
)

func HandleSignup(w http.ResponseWriter, r *http.Request, db storage.StorageInterface) {
	// Get user data from request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "User with bad format", serverErrors.GenericError)
		return
	}
	// Verify if user already exists
	userMatch, err := db.GetUserByEmail(user.Email)
	if userMatch != nil {
		utils.SendError(w, http.StatusNotAcceptable, "User already exists with this email", serverErrors.UserAlreadyExistsForEmail)
		return
	}
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to validate input data", serverErrors.GenericError)
		return
	}
	// Encript user password
	encriptedPassword, err := utils.GenerateHashPassword(user.Password)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to encript password", serverErrors.GenericError)
		return
	}
	user.Password = encriptedPassword
	// Add default user role
	user.Role = "standard"
	// Create user in database
	err = db.CreateUser(&user)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create user", serverErrors.GenericError)
		return
	}
	// Create JWT
	tokenString, err := utils.CreateJWTToken(user.Role, user.Email)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create JWT", serverErrors.GenericError)
		log.Printf("error: %v", err)
		return
	}
	// Send response
	w.Header().Set("Bearer-Token", tokenString)
	utils.SendJSON(w, http.StatusCreated, nil)
}

func HandleLogin(w http.ResponseWriter, r *http.Request, db storage.StorageInterface) {
	// Get user data from request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "User with bad format", serverErrors.GenericError)
		return
	}
	// Check if user exists on database using email to find and password to validate user
	userMatch, err := db.GetUserByEmail(user.Email)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to validate input data", serverErrors.GenericError)
		return
	}
	if userMatch == nil {
		utils.SendError(w, http.StatusNotFound, "User does not exist", serverErrors.InvalidEmailOrPassword)
		return
	}
	isValidPassWord := utils.CompareHashPassword(user.Password, userMatch.Password)
	if !isValidPassWord {
		utils.SendError(w, http.StatusForbidden, "Wrong password", serverErrors.InvalidEmailOrPassword)
		return
	}
	// Create JWT
	tokenString, err := utils.CreateJWTToken(userMatch.Role, userMatch.Email)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create JWT", serverErrors.GenericError)
		log.Printf("error: %v", err)
		return
	}
	// Send response
	w.Header().Set("Bearer-Token", tokenString)
	utils.SendJSON(w, http.StatusOK, nil)
}

func HandleRefreshToken(w http.ResponseWriter, r *http.Request, db storage.StorageInterface) {
	// Get user data from request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "User with bad format", serverErrors.GenericError)
		return
	}
	if user.Email == "" || user.Password == "" {
		utils.SendError(w, http.StatusBadRequest, "User with bad format", serverErrors.GenericError)
		return
	}
	// Check if user exists on database using email to find and password to validate user
	userMatch, err := db.GetUserByEmail(user.Email)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to validate input data", serverErrors.GenericError)
		return
	}
	if userMatch == nil {
		utils.SendError(w, http.StatusNotFound, "Failed to validate input data", serverErrors.InvalidEmailOrPassword)
		return
	}
	isValidPassWord := utils.CompareHashPassword(user.Password, userMatch.Password)
	if !isValidPassWord {
		utils.SendError(w, http.StatusForbidden, "Failed to validate input data", serverErrors.InvalidEmailOrPassword)
		return
	}
	// Create JWT
	tokenString, err := utils.CreateJWTToken(userMatch.Role, userMatch.Email)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create JWT", serverErrors.GenericError)
		log.Printf("error: %v", err)
		return
	}
	// Send response
	w.Header().Set("Bearer-Token", tokenString)
	utils.SendJSON(w, http.StatusOK, nil)
}
