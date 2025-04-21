package initialize

import (
	"fmt"

	"github.com/nk-hung/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {

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

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
