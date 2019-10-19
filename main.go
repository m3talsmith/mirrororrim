package main

import (
	"github.com/spf13/viper"
	"github.com/m3talsmith/mirrororrim/cmd"
)

func init() {
	// Setup cli flags
	
	// Setup the config structure
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.mirrororrim")
	viper.SetConfigType("toml")
	viper.SafeWriteConfig()
}

func main() {
	cmd.Exec()
}