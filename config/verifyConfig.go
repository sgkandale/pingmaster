package config

import "log"

func GetVerifiedConfig() config {
	parsedConfig := ParseConfig()

	// Server Checks
	if parsedConfig.Server.Port == 0 {
		log.Fatal("[ERROR] server port is not set")
	}
	if parsedConfig.Server.Port < 0 {
		log.Fatal("[ERROR] server port is invalid")
	}
	if parsedConfig.Server.TLS {
		if parsedConfig.Server.CertPath == "" {
			log.Fatal("[ERROR] server tls cert path is not set")
		}
		if parsedConfig.Server.KeyPath == "" {
			log.Fatal("[ERROR] server tls key path is not set")
		}
	}
	if parsedConfig.Server.PathPrefix != "" {
		if parsedConfig.Server.PathPrefix[0:1] != "/" {
			log.Fatal("[ERROR] server path prefix must start with '/'")
		}
		if parsedConfig.Server.PathPrefix[len(parsedConfig.Server.PathPrefix)-1:] == "/" {
			log.Fatal("[ERROR] server path prefix must not end with '/'")
		}
	}

	// database checks
	if parsedConfig.Database.DatabaseType == "" {
		log.Fatal("[ERROR] database type is not set")
	}
	if parsedConfig.Database.Host == "" {
		log.Fatal("[ERROR] database host is not set")
	}
	if parsedConfig.Database.Port <= 0 {
		log.Fatal("[ERROR] database port is not set")
	}
	if parsedConfig.Database.DatabaseName == "" {
		log.Fatal("[ERROR] database name is not set")
	}
	if parsedConfig.Database.TimeoutInSeconds <= 0 {
		log.Fatal("[ERROR] database timeout in seconds is not set")
	}

	return parsedConfig
}
