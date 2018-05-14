package main

import (
	"os"
)

// GetEnv is a helper function used to provide a default value if the environment var doesn't exist
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
