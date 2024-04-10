package internal

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func ConfigInit() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))

	fmt.Println("mysql ip: ", viper.GetString("mysql.host"))
	fmt.Println("mysql port: ", viper.GetInt("mysql.port"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())
}
