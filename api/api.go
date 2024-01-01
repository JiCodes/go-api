package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Succcess code, 200 for success
	Code int 
	// Account Balance
	Balance float64
}

// Error Response
type Error struct {
	// Error Code
	Code int
	// Error Message
	Message string
}

