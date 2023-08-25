package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func getAccessToken() string {
	url := fmt.Sprintf(WeChatUrl, WeChatAppid, WeChatSecret)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("获取微信token失败", err)
		return ""
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("获取微信accesstoken失败 结束请求：", err.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("微信token读取失败", err)
		return ""
	}

	token := token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("微信token解析json失败", err)
		return ""
	}

	return token.AccessToken
}

//发送模板消息
func templatepost(access_token string, reqdata string, fxurl string, templateid string, openid string) error {
	fmt.Println("------------------------sending-------------------------")
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + access_token

	reqbody := "{\"touser\":\"" + openid + "\", \"template_id\":\"" + templateid + "\", \"url\":\"" + fxurl + "\", \"data\": " + reqdata + "}"
	fmt.Println(strings.NewReader(string(reqbody)))
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(string(reqbody)))

	if err != nil {
		fmt.Println("--------------------post err--------------------")
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("-----------read err-------------")
		fmt.Println(err)
		return err
	}

	fmt.Println(string(body))
	return nil
}

type subList struct {
	Count   uint   `json:"count"` //关注人数
	OpenIds IdList `json:"data"`
	nextId  string `json:"next_openid"`
	total   uint   `json:"total"` //
}

type IdList struct {
	Openid []string `json:"openid"`
}

//获取关注者列表
func getflist(access_token string) subList {
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + access_token + "&next_openid="
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取关注列表失败", err)
		return subList{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容失败", err)
		return subList{}
	}

	var data subList //map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("---------------------序列号列表失败----------------------")
		fmt.Println(err)
		return subList{}
	}
	fmt.Println(data)
	return data
}
