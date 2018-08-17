package env

import (
	"os"
)

func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)

	if true == exists {
		return value
	}

	return fallback
}
