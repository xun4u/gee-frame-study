package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {

	//实例化框架
	r := gee.New()
	//添加路由
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
	})
	r.GET("hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	//运行HTTP服务
	r.Run(":9999")
}
