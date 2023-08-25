package main

const (
	WeatherAppid     = ""
	WeatherAppsecret = ""
	WeatherApi       = "https://www.yiketianqi.com/free/week?unescape=1&appid=%s&appsecret=%s&city=%s"
)

const (
	WeChatUrl    = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v"
	WeChatAppid  = ""
	WeChatSecret = ""
	WechatTemp   = ""
)

// 发送对象的openID
var sendList = []string{
	"",
}
