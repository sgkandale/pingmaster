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

	return parsedConfig
}
