package utils

import (
	"encoding/json"
	"net/http"
)

// Message ...
func FormatResponse(statusCode int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"status": statusCode, "message": message, "data": data}
}

// Respond ...
func SendResponse(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
