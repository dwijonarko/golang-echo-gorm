package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Config {
	configuration := Config{}
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Println(err)
	}

	return configuration
}
