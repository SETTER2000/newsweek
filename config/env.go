package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	log.Println(".env file loaded")
}

type DatabaseConfig struct {
	url string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: os.Getenv("DATABASE_URL"),
	}
}

func getString(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func getInt(key string, defaultValue int) int {
	v := os.Getenv(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(k string, defaultValue bool) bool {
	v := os.Getenv(k)
	b, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}
	return b
}
