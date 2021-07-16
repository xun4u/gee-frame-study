package main

import (
	"gee"
	"net/http"
)

func main() {

	//实例化框架
	r := gee.New()
	//添加路由
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,path:%s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	//运行HTTP服务
	r.Run(":9999")
}
