package web

import (
	"encoding/json"
	"net/http"
)

func GenerateErrorResponse(w http.ResponseWriter, message, description string, statusCode int) {
	json.NewEncoder(w).Encode(struct {
		Message     string `json:"message"`
		Description string `json:"description"`
		Status      int    `json:"status"`
	}{
		Message:     message,
		Description: description,
		Status:      statusCode,
	})
}

func GenerateDefaultResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	json.NewEncoder(w).Encode(struct {
		Data interface{} `json:"data"`
	}{
		Data: data,
	})
}
