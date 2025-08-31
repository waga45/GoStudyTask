package study

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func GetRequest() {
	var url = "http://www.baidu.com"

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

type UserInfo struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func TestLogin() {
	url := "https://www.dawei.art/dwart/user/user/userLogin"
	var payload = []byte(`{"username":"ddd","password":"111"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "application/json")

	client := http.Client{Timeout: time.Second * 100}

	respon, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(respon)
	resultBuffer, err := io.ReadAll(respon.Body)
	defer respon.Body.Close()
	if err != nil {
		fmt.Println("返回数据读取失败")
		fmt.Println(err)
		return
	}
	fmt.Println(string(resultBuffer))
	userInfo := UserInfo{}
	var e = json.Unmarshal(resultBuffer, &userInfo)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("返回数据：", userInfo)
}
