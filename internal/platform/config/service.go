package config

import (
	"os"
	"time"
)

var configService *Config

func InitConfigService() {
	configService = &Config{
		ServerConfig: ServerConfig{
			AppName:     getEnv("APP_NAME", "alutec-inventory-be"),
			Port:        getEnv("PORT", "8080"),
			GinMode:     getEnv("GIN_MODE", "debug"),
			FrontendURL: getEnv("FRONTEND_URL", "http://localhost:9000"),
		},
		DatabaseConfig: DatabaseConfig{
			DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:54324/alutec-inventory-db?sslmode=disable"),
		},
		AuthConfig: AuthConfig{
			Email:    getEnv("AUTH_EMAIL", ""),
			Password: getEnv("AUTH_PASSWORD", ""),
			Secret:   getEnv("AUTH_SECRET", ""),
			TokenTTL: 12 * time.Hour,
		},
	}
}

func GetConfigService() *Config {
	return configService
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
