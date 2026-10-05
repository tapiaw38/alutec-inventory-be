package config

import "time"

type (
	ServerConfig struct {
		AppName     string
		Port        string
		GinMode     string
		FrontendURL string
	}

	DatabaseConfig struct {
		DatabaseURL string
	}

	// AuthConfig holds the single set of credentials the API accepts. It is a
	// placeholder for real user accounts, hence one email and one password.
	AuthConfig struct {
		Email    string
		Password string
		Secret   string
		TokenTTL time.Duration
	}

	Config struct {
		ServerConfig   ServerConfig
		DatabaseConfig DatabaseConfig
		AuthConfig     AuthConfig
	}
)
