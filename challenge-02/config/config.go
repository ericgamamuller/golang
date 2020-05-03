package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Initialize() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	path, _ := os.Getwd()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal Error while reading config.yaml: %s \n", err))
	}
}

func ReadString(field string) string {
	return viper.GetString(field)
}
