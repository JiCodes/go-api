package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/JiCodes/go-api/internal/handlers/middleware"
)

func Handlers(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes) // strip slashes from incoming requests

	r.Route("/account", func(router chi.Router) {
		// middleware for /account route
		router.Use(middleware.AuthMiddleware) 
		
		router.Get("/coins", GetCoinBalance)
	})

}