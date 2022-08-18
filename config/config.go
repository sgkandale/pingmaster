package config

type ServerConfig struct {
	Port     int
	TLS      bool
	CertPath string
	KeyPath  string
}

type config struct {
	Server ServerConfig
}
