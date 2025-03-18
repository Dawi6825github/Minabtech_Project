package utils

import (
    "encoding/json"
    "net/http"
)

// ResponseSuccess sends a success message with a response payload
func ResponseSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := map[string]interface{}{
        "status":  "success",
        "message": message,
        "data":    data,
    }
    json.NewEncoder(w).Encode(response)
}

// ResponseError sends an error message with an error code
func ResponseError(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    response := map[string]interface{}{
        "status":  "error",
        "message": message,
    }
    json.NewEncoder(w).Encode(response)
}
