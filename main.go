package main

import (
	"fmt"
	"life/actions"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"

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

	c := cron.New()
	c.AddFunc("0 10-20 * * 1-5", actions.DrinkWater)
	c.AddFunc("0 10,18 * * *", actions.SendWeather)
	//	c.AddFunc("1-59 * * * *", func() { fmt.Println("helo") })
	c.Start()

	//阻止main推出
	ch := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	//阻塞直到有信号传入
	fmt.Println("启动")
	//阻塞直至有信号传入
	s := <-ch
	fmt.Println("退出信号", s)
}
