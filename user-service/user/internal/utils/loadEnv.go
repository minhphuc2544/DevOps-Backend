package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	paths := []string{
		".env",
		"../.env",
		"../../.env",
	}

	for _, path := range paths {
		err := godotenv.Load(path)
		if err == nil {
			log.Printf("Loaded .env from: %s", path)
			return nil
		}
	}

	return fmt.Errorf("failed to load .env from known paths")
}
