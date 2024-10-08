package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/1l1k3tw3nty5/rssfeed/auth"
	"github.com/1l1k3tw3nty5/rssfeed/internal/database"
	"github.com/google/uuid"
)

func (apiConf *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Failed to parse JSON into parameters: %v", err)
		respondWithError(w, 400, fmt.Sprintf("Failed to parse JSON: %v", err))
		return
	}

	user, err := apiConf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Printf("Failed to create new user: %v", err)
		respondWithError(w, 500, fmt.Sprintln("Internal Error"))
	}
	respondWithJson(w, 201, databaseUserToUser(user))

}

func (apiConf *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		log.Println("Failed to get api key: ", err)
		respondWithError(w, 403, fmt.Sprintf("Auth failed: %v", err))
		return
	}

	userByApiKey, err := apiConf.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		log.Printf("Failed to get user by api key: %v", err)
		respondWithError(w, 400, fmt.Sprintf("Failed to get user: %v", err))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(userByApiKey))
}
