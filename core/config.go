package core

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Debug bool
	Port  string
}

var config Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	viper.SetDefault("debug", true)
	viper.SetDefault("port", ":8080")

	config.Debug = viper.GetBool("debug")
	config.Port = viper.GetString("port")
}

func GetConfig() *Config {
	return &config
}
