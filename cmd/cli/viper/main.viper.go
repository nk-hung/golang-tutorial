package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Using config PORT file:", viper.GetString("server.port"))
	fmt.Println("Using config HOST file:", viper.GetString("server.host"))

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Using config file:", config.Server.Port)
	fmt.Println("Using config file:", config.Databases[0].Port)
}
