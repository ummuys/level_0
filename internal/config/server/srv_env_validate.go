package config

import (
	"fmt"
	"os"
)

func ParseServerEnv() (string, error) {
	port := os.Getenv("APP_PORT")
	if port == "" {
		return "", fmt.Errorf("invalid app_port")
	}
	return port, nil
}
