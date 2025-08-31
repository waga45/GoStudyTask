package study

import (
	"net/http"
)

func authHandle(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("CustomeKey", "张三")
	d := request.URL.Query()
	var name = d.Get("username")
	var pwd = d.Get("pwd")
	if name == "张三" && pwd == "123" {
		response.Write([]byte("登入成功"))
	} else {
		response.Write([]byte("登入成失败"))
	}
}

func StartWebServer() {
	http.HandleFunc("/login", authHandle)

	http.ListenAndServe(":8089", nil)
}
