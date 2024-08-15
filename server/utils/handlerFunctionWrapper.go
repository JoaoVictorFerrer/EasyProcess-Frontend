package utils

import (
	"EasyProcess/serverErrors"
	"EasyProcess/storage"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type serverFuncWithMongoClient func(w http.ResponseWriter, r *http.Request, db storage.StorageInterface)
type standardHandlerFunction func(w http.ResponseWriter, r *http.Request)

func HTTPHandleFunc(myServerFunc serverFuncWithMongoClient, db storage.StorageInterface) standardHandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		myServerFunc(w, r, db)
	}
}

func ProtectedHTTPHandleFunc(myServerFunc serverFuncWithMongoClient, db storage.StorageInterface) standardHandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from header
		token := r.Header.Get("Bearer-Token")
		// If token is empty send no auth resonse and return
		if token == "" {
			SendError(w, http.StatusForbidden, "User not authenticated", serverErrors.GenericError)
			return
		}
		// Token is not empty validate it
		_, err := ValidateJWTToken(token)
		if err != nil && err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			// Token is expired
			SendError(w, http.StatusForbidden, "User not authenticated", serverErrors.ExpiredToken)
			return
		}
		if err != nil {
			SendError(w, http.StatusForbidden, "Failed to atenticate user", serverErrors.GenericError)
			return
		}
		// User is authenticated pass request forward
		myServerFunc(w, r, db)
	}
}
