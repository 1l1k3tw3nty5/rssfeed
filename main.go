package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("PORT is not found")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Ling"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routerV1 := chi.NewRouter()

	routerV1.Get("/alive", handlerAlive)
	routerV1.Get("/error", handlerErr)

	router.Mount("/v1", routerV1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	log.Printf("Server is running on port %v", portNumber)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
