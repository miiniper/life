package main

import (
	"fmt"
	"life/actions"

	"github.com/fsnotify/fsnotify"
	"github.com/miiniper/loges"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	fmt.Println("starting ... ")

	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		loges.Loges.Info("Config file changed: ", zap.Any("", e.Name))
	})

	actions.SendWeather()
	//actions.DrinkWater()

}
