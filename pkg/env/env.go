package env

import (
	"os"
)

// Environment ...
func Environment() string {
	return GetFromEnvOrDefault("ENVIRONMENT", "local")
}

// AWSRegion ...
func AWSRegion() string {
	return GetFromEnvOrDefault("AWS_REGION", "ap-southeast-2")
}

// Account ...
func Account() string {
	return GetFromEnvOrDefault("DB_LIST", "db-list")
}

// GetFromEnvOrDefault - returns the environment variable if it exists else returns to default
func GetFromEnvOrDefault(name, defaultValue string) string {
	fromEnv := os.Getenv(name)
	if fromEnv == "" {
		return defaultValue
	}
	return fromEnv
}