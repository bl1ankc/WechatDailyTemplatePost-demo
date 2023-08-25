package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	City   string         `json:"city"`
	CityId string         `json:"cityid"`
	Data   []dailyWeather `json:"data"`
}
type dailyWeather struct {
	Date     string `json:"date"`      //日期
	TemDay   string `json:"tem_day"`   //白天温度
	TemNight string `json:"tem_night"` //晚上温度
	Wea      string `json:"wea"`       //天气
	Win      string `json:"win"`       //风向
	WinSpeed string `json:"win_speed"` //风级
}

func getWeather(city string) Weather {
	url := fmt.Sprintf(WeatherApi, WeatherAppid, WeatherAppsecret, city)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Get weather worry", err)
		return Weather{}
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read information failed", err)
		return Weather{}
	}

	var data Weather
	err = json.Unmarshal(body, &data)
	fmt.Println(data)
	return data
}
