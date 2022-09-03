package config

import (
	"log"

	"github.com/spf13/viper"
)

func ParseConfig() Config {
	log.Println("[INF] parsing config file")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("[ERR] parsing config file : ", err)
	}

	readConfig := Config{}

	err = viper.Unmarshal(&readConfig)
	if err != nil {
		log.Fatal("[ERR] unmarshing config file in struct : ", err)
	}

	log.Println("[INF] config file parsed successfully")

	return readConfig
}
