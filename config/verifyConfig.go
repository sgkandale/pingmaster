package config

import "log"

func GetVerifiedConfig() Config {
	parsedConfig := ParseConfig()

	// Server Checks
	if parsedConfig.Server.Port == 0 {
		log.Fatal("[ERR] server port is not set")
	}
	if parsedConfig.Server.Port < 0 {
		log.Fatal("[ERR] server port is invalid")
	}
	if parsedConfig.Server.TLS {
		if parsedConfig.Server.CertPath == "" {
			log.Fatal("[ERR] server tls cert path is not set")
		}
		if parsedConfig.Server.KeyPath == "" {
			log.Fatal("[ERR] server tls key path is not set")
		}
	}
	if parsedConfig.Server.PathPrefix != "" {
		if parsedConfig.Server.PathPrefix[0:1] != "/" {
			log.Fatal("[ERR] server path prefix must start with '/'")
		}
		if parsedConfig.Server.PathPrefix[len(parsedConfig.Server.PathPrefix)-1:] == "/" {
			log.Fatal("[ERR] server path prefix must not end with '/'")
		}
	}

	// database checks
	if parsedConfig.Database.DatabaseType == "" {
		log.Fatal("[ERR] database type is not set")
	}
	if parsedConfig.Database.Host == "" {
		log.Fatal("[ERR] database host is not set")
	}
	if parsedConfig.Database.Port <= 0 {
		log.Fatal("[ERR] database port is not set")
	}
	if parsedConfig.Database.DatabaseName == "" {
		log.Fatal("[ERR] database name is not set")
	}
	if parsedConfig.Database.TimeoutInSeconds <= 0 {
		log.Fatal("[ERR] database timeout in seconds is not set")
	}
	if parsedConfig.Database.MaxConcurrentQueries == 0 {
		log.Printf("[WRN] database max concurrent queries not set, using default value %d", Default_DBMaxConcurrentQuries)
		parsedConfig.Database.MaxConcurrentQueries = Default_DBMaxConcurrentQuries
	}

	// security checks
	if parsedConfig.Security.TokenSecret == "" {
		log.Fatal("[ERR] security token secret is not set")
	}
	if len(parsedConfig.Security.TokenSecret) < 12 {
		log.Fatal("[ERR] security token secret should be atleast 12 characters long")
	}
	if len(parsedConfig.Security.TokenSecret) > 64 {
		log.Fatal("[ERR] security token secret should ot be longer than 64 characters")
	}
	if parsedConfig.Security.AllowedOrigins == "" {
		log.Fatal("[ERR] security allowed origin is not set")
	}

	return parsedConfig
}
