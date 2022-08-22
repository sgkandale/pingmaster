package config

import (
	"log"

	"github.com/spf13/viper"
)

func ParseConfig() config {
	log.Println("[INFO] parsing config file")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("[ERROR] parsing config file : ", err)
	}

	readConfig := config{}

	err = viper.Unmarshal(&readConfig)
	if err != nil {
		log.Fatal("[ERROR] unmarshing config file in struct : ", err)
	}

	log.Println("[INFO] config file parsed successfully")

	return readConfig
}
