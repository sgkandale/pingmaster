package config

type ServerConfig struct {
	Port       int
	TLS        bool
	CertPath   string
	KeyPath    string
	PathPrefix string
}

type config struct {
	Server ServerConfig
}
