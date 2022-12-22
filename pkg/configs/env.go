package configs

import (
	"fmt"
	"os"
)

type Environment struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

func GetEnv(value string) (string, error) {
	_, isPresent := os.LookupEnv(value)
	if !isPresent {
		return "", fmt.Errorf("%s not found in environment", value)
	}
	return os.Getenv(value), nil
}
