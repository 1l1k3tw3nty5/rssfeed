package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		log.Println("No authorization data provided")
		return "", fmt.Errorf("wrong authorization data provided")
	}

	apiKey := strings.Split(val, "=")
	if length := len(apiKey); length != 2 {
		log.Printf("Wrong api key format (length): %v", length)
		return "", fmt.Errorf("wrong api key format")
	}

	if apiKey[0] != "api_key" {
		log.Printf("Wrong api key format (keyword): %v", apiKey[0])
		return "", fmt.Errorf("wrong api key format")
	}

	return apiKey[1], nil
}
