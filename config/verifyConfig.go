package config

import "log"

func GetVerifiedConfig() config {
	parsedConfig := ParseConfig()

	// Server Checks
	if parsedConfig.Server.Port == 0 {
		log.Fatal("server port is not set")
	}
	if parsedConfig.Server.Port < 0 {
		log.Fatal("server port is invalid")
	}

	return parsedConfig
}
