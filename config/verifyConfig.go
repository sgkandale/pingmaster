package config

import "log"

func GetVerifiedConfig() Config {
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

	// other checks
	if parsedConfig.TokenSecret == "" {
		log.Fatal("[ERROR] token secret is not set")
	}
	if len(parsedConfig.TokenSecret) < 12 {
		log.Fatal("[ERROR] token secret should be atleast 12 characters long")
	}
	if len(parsedConfig.TokenSecret) > 64 {
		log.Fatal("[ERROR] token secret should ot be longer than 64 characters")
	}

	return parsedConfig
}
