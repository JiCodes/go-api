package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Succcess code, 200 for success
	Code int 
	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error Code
	Code int
	// Error Message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	// create error response
	errorResponse := Error{
		Code: code,
		Message: message,
	}

	// set response header
	w.Header().Set("Content-Type", "application/json")
	// set response status code
	w.WriteHeader(code)
	// write error response
	json.NewEncoder(w).Encode(errorResponse)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Occurred", http.StatusInternalServerError)
	}
)