package initialize

import (
	"fmt"
	"learning-french-service/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	fmt.Println(viper.Get("server.port"))

	// config struct
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
