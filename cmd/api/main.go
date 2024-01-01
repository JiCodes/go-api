package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" // web framework
	log "github.com/sirupsen/logrus" // set log as alias for logrus
	"github.com/JiCodes/go-api/internal/handlers" 
)

func main() {
	log.SetReportCaller(true) // set log to report caller

	var r *chi.Mux = chi.NewRouter() // create new router

	handlers.Handlers(r) // call handlers function

	fmt.Println("Starting Go API service ...")

	err := http.ListenAndServe("localhost:8080", r) // start server

	if err != nil {
		log.Error(err) // log the error
	}

}
