package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func StartServer(router *http.ServeMux) {
	port := "8080" //os.Getenv("PORT")
	if os.Getenv("ENV") == "DEV" {
		// We are running locally, use local port value from .env file
		port = os.Getenv("LOCAL_HOST_PORT")
	}
	middlewareStack := CreateMiddlewareStack(
		LogIncommingRequestInfo,
	)
	handlerWithCors := cors.AllowAll().Handler(middlewareStack(router))
	log.Printf("🌍 Server listening on port: %v 🌍", port)
	panic(http.ListenAndServe("0.0.0.0:"+port, handlerWithCors))
}
