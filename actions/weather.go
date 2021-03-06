package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"life/unit"
	"os"

	"github.com/spf13/viper"

	"github.com/miiniper/loges"
	"go.uber.org/zap"
)

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

var GaoDeToken string
var BotToken string

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

//330106 西湖区
func GetWeatherForecasts(adCode string) WeatherForecasts {
	weatherUrl := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s&extensions=all", GaoDeToken, adCode)
	resp := unit.GetUrl(weatherUrl)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		loges.Loges.Info("http code not 200 ,", zap.Int("status code", resp.StatusCode))
	}

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
	bot := NewLifeBot("lifeBot", BotToken)
	w := GetWeather("330106")

	msg := FmtWeather(w)
	bot.SendMsg("911000205", msg)
}

func FmtWeather(w WeatherInfo) string {
	msg := fmt.Sprintf("当前天气情况:\n 位置：%s\n 天气：%s\n 风向：%s\n 风力：%s\n 温度：%s\n 湿度：%s\n", w.Lives[0].City, w.Lives[0].Weather, w.Lives[0].WindDirection, w.Lives[0].WindPower, w.Lives[0].Temperature, w.Lives[0].Humidity)
	return msg
}
