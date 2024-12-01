package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins (no credentials)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false, // Set to false when using "*"
		MaxAge:           300,
	}).Handler)
	
	v1Router := chi.NewRouter() // create a new router for the /v1 API

	v1Router.Get("/healthz", handlerReadiness) // the /ready endpoint is handled by the handlerReadiness function
	v1Router.Get("/err", handlerError) // the /ready endpoint is handled by the handlerReadiness function

	router.Mount("/v1", v1Router) // mount the /v1 API router under the root router

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server listening on port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}