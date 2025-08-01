package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println(viper.Get("server.port"))

	// config struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	for _, db := range config.Database {
		fmt.Println(db)
	}
}
