package config

type ServerConfig struct {
	Port       int
	TLS        bool
	CertPath   string
	KeyPath    string
	PathPrefix string
}

type SecurityConfig struct {
	TokenSecret    string
	AllowedOrigins string
}

type DatabaseConfig struct {
	DatabaseType     string
	Username         string
	Password         string
	Host             string
	Port             int
	DatabaseName     string
	TimeoutInSeconds int
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Security SecurityConfig
}
