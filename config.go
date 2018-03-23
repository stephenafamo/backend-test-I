package main

import (
	"github.com/spf13/viper"
)

func setupConfig() {
	setDefaults()
	loadConfig()
}

func setDefaults() {
	viper.SetDefault("SheetIndex", 0)
	viper.SetDefault("defaultMinFollowers", 1000)
	viper.SetDefault("defaultMaxFollowers", 50000)
}

func loadConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // type of config file
	viper.AddConfigPath(".")      // look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	checkError(err)               // Handle errors reading the config file
}
