package utils

import (
	"EasyProcess/serverErrors"
	"encoding/json"
	"net/http"
)

type ServerResponse struct {
	Data              any    `json:"data"`
	Error             bool   `json:"error"`
	ErrorText         string `json:"errorText"`
	InternalErrorCode int    `json:"internalErrorCode"`
}

func SendJSON(w http.ResponseWriter, status int, object any) {
	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "Bearer-Token")
	w.WriteHeader(status)
	// Build response
	serverResponse := ServerResponse{
		Data:              object,
		Error:             false,
		InternalErrorCode: serverErrors.NoError,
	}
	// Encode and send JSON
	json.NewEncoder(w).Encode(serverResponse)
}

func SendRawJSON(w http.ResponseWriter, status int, object any) {
	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "Bearer-Token")
	w.WriteHeader(status)
	// Encode and send JSON
	json.NewEncoder(w).Encode(object)
}

func SendError(w http.ResponseWriter, status int, errorText string, internalErrorCode int) {
	// Setup header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "Bearer-Token")
	w.WriteHeader(status)
	// Build response
	serverResponse := ServerResponse{
		Error:             true,
		ErrorText:         errorText,
		InternalErrorCode: internalErrorCode,
	}
	// Encode and send JSON
	json.NewEncoder(w).Encode(serverResponse)
}
