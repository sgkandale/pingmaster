package config

const (
	Default_DBMaxConcurrentQuries int = 100
	Default_DBPingsValidity       int = 30
)

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
	DatabaseType         string
	Username             string
	Password             string
	Host                 string
	Port                 int
	DatabaseName         string
	TimeoutInSeconds     int
	MaxConcurrentQueries int
	PingsValidity        int
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Security SecurityConfig
}
