package config

type ServerConfig struct {
	Port       int
	TLS        bool
	CertPath   string
	KeyPath    string
	PathPrefix string
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

type config struct {
	Server   ServerConfig
	Database DatabaseConfig
}
