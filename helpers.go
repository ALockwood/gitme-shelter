package main

import (
	"os"
	"strings"
)

//Check if a string is nil or empty after trimming all whitespace
func stringIsNilOrEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
