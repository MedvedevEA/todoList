package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BindAddress           string
	DatabaseConnectString string
}

func New() (*Config, error) {
	if err := godotenv.Load("./configs/config.env"); err != nil {
		log.Printf("Сonfiguration file error: %s", err.Error())
	}
	return &Config{
		BindAddress:           getEnv("TODOLIST_SERVER_BIND_ADDRESS", ":8000"),
		DatabaseConnectString: getEnv("TODOLIST_DATABASE_CONNECT_STRING", "host=localhost database=postgres port=5432 sslmode=disable user=postgres password=postgres"),
	}, nil
}
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
