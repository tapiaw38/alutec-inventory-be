package config

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

	Config struct {
		ServerConfig   ServerConfig
		DatabaseConfig DatabaseConfig
	}
)
