package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"life/unit"
	"os"

	bot2 "github.com/miiniper/tgmsg_bot/bot"

	"github.com/fsnotify/fsnotify"
	"github.com/miiniper/loges"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var BotToken string
var GaoDeToken string

func init() {
	BotToken = os.Getenv("BotToken")
	if BotToken == "" {
		BotToken = viper.GetString("BotToken")
	}

	GaoDeToken = os.Getenv("GaoDeToken")
	if GaoDeToken == "" {
		GaoDeToken = viper.GetString("GaoDeToken")
	}
	if BotToken == "" || GaoDeToken == "" {
		panic("BotToken or GaoDeToken is null, start server fail....")
	}

}

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
		//fmt.Printf("Config file changed: %s", e.Name)
		loges.Loges.Info("Config file changed: ", zap.Any("", e.Name))
	})

	SendWeather()

}

type HttpWeather struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	InfoCode string `json:"infoCode"`
}

type Forecast struct {
	City       string `json:"city"`
	Adcode     string `json:"adCode"`
	Province   string `json:"province"`
	ReportTime string `json:"reportTime"`
	Casts      []Cast `json:"casts"`
}
type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	DayWeather   string `json:"dayWeather"`
	NightWeather string `json:"nightWeather"`
	DayTemp      string `json:"dayTemp"`
	NightTemp    string `json:"nightTemp"`
	DayWind      string `json:"dayWind"`
	NightWind    string `json:"nightWind"`
	DayPower     string `json:"dayPower"`
	NightPower   string `json:"nightPower"`
}

type Live struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	AdCode        string `json:"adCode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	WindDirection string `json:"windDirection"`
	WindPower     string `json:"windPower"`
	Humidity      string `json:"humidity"`
	ReportTime    string `json:"reportTime"`
}

type WeatherForecasts struct {
	HttpWeather HttpWeather `json:"httpWeather"`
	Forecasts   []Forecast  `json:"forecasts"`
}

type WeatherInfo struct {
	HttpWeather HttpWeather `json:"httpWeather"`
	Lives       []Live      `json:"lives"`
}

//330106 西湖区
func GetWeatherForecasts(adCode string) WeatherForecasts {
	weatherUrl := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s&extensions=all", GaoDeToken, adCode)
	resp := unit.GetUrl(weatherUrl)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		loges.Loges.Info("http code not 200 ,", zap.Int("status code", resp.StatusCode))
	}
	//loges.Loges.Info("", zap.String("body", string(body)))
	//	fmt.Println(string(body))
	var w WeatherForecasts
	err := json.Unmarshal(body, &w)
	if err != nil {
		loges.Loges.Error("fmt json error ", zap.Error(err))
	}
	return w
}

func GetWeather(adCode string) WeatherInfo {
	weatherUrl := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s", GaoDeToken, adCode)
	resp := unit.GetUrl(weatherUrl)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		loges.Loges.Info("http code not 200 ,", zap.Int("status code", resp.StatusCode))
	}
	//	fmt.Println(string(body))
	var w WeatherInfo
	err := json.Unmarshal(body, &w)
	if err != nil {
		loges.Loges.Error("fmt json error ", zap.Error(err))
	}
	return w
}

func SendWeather() {
	bot, err := bot2.NewBotApi(BotToken, "lifeBot")
	if err != nil {
		loges.Loges.Error("get bot error", zap.Error(err))
	}
	w := GetWeather("330106")

	msg := fmt.Sprintf("位置：%s\n 天气：%s\n 风向：%s\n 风力：%s\n 温度：%s\n", w.Lives[0].City, w.Lives[0].Weather, w.Lives[0].WindDirection, w.Lives[0].WindPower, w.Lives[0].Temperature)
	bot.SendMsg("911000205", msg)
}
