package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigStruct struct {
	Logger struct {
		Level   string   `mapstructure:"level"`
		Outputs []string `mapstructure:"outputs"`
	} `mapstructure:"logger"`
}

func GetConfig() ConfigStruct {

	viper.SetConfigName("config")    // name of config file (without extension)
	viper.AddConfigPath("./configs") // path to look for the config file in
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	var config ConfigStruct

	err = viper.Unmarshal(&config)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("Can't unmarshal: %s \n", err)
	}
	return config
}
