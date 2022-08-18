package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var (
	configFilePath string
)

func init() {
	configFile := flag.String("config", ".", "path to config file")

	flag.Parse()

	configFilePath = *configFile
}

func ParseConfig() config {
	log.Println("[INFO] parsing config file")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)
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
