package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/1l1k3tw3nty5/rssfeed/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		log.Fatal("PORT is not found")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database connection string is not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unable to open a connection with the database")
	}

	apiConf := apiConfig{
		DB: database.New(conn),
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
	routerV1.Post("/user", apiConf.handlerCreateUser)

	router.Mount("/v1", routerV1)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portNumber,
	}

	log.Printf("Server is running on port %v", portNumber)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
