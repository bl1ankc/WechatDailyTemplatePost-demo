package main

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	firstDay = "2022-09-26"
)

type Location struct {
	Value string `json:"value"`
}
type Wea struct {
	Value string `json:"value"`
}
type TemDay struct {
	Value string `json:"value"`
}
type TemNight struct {
	Value string `json:"value"`
}
type Win struct {
	Value string `json:"value"`
}
type WinSpeed struct {
	Value string `json:"value"`
}
type LovesDay struct {
	Value int64 `json:"value"`
}
type TomorrowTD struct {
	Value string `json:"value"`
}
type TomorrowTN struct {
	Value string `json:"value"`
}
type TomorrowWea struct {
	Value string `json:"value"`
}
type Content struct {
	Location    `json:"location"`   //地址
	Wea         `json:"wea"`        //天气
	TemDay      `json:"temDay"`     //白天气温
	TemNight    `json:"temNight"`   //晚上气温
	Win         `json:"win"`        //风向
	WinSpeed    `json:"winSpeed"`   //风力
	TomorrowTD  `json:"tomorrowTD"` //明日白天气温
	TomorrowTN  `json:"tomorrowTN"` //每日晚上
	TomorrowWea `json:"tomorrowWea"`
	LovesDay    `json:"lovesDay"`
}

func SendMessage(accessToken string, openid string) {
	//l day
	f, _ := time.Parse("2006-01-02", firstDay)
	d := time.Now().Local().Sub(f)
	day := int64(d.Hours() / 24)

	//struct
	weather := getWeather("深圳")
	data := Content{
		Location:    Location{"深圳"},
		Wea:         Wea{weather.Data[0].Wea},
		TemDay:      TemDay{weather.Data[0].TemDay},
		TemNight:    TemNight{weather.Data[0].TemNight},
		Win:         Win{weather.Data[0].Win},
		WinSpeed:    WinSpeed{weather.Data[0].WinSpeed},
		TomorrowTN:  TomorrowTN{weather.Data[1].TemDay},
		TomorrowTD:  TomorrowTD{weather.Data[1].TemNight},
		TomorrowWea: TomorrowWea{weather.Data[1].Wea},
		LovesDay:    LovesDay{day},
	}

	reqData, _ := json.Marshal(data)
	fmt.Println("reaData=", string(reqData))

	err := templatepost(accessToken, string(reqData), "", WechatTemp, openid)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("send access")
	}

}

func Send() {
	accessToken := getAccessToken()
	ssList := getflist(accessToken)

	for _, v := range ssList.OpenIds.Openid {
		SendMessage(accessToken, v)
		//for _, u := range sendList {
		//	if v == u {
		//
		//	}
		//}
	}
}
