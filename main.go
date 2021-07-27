package main

import (
	"gee"
	"net/http"
)

func main() {

	//实例化框架
	r := gee.New()

	//全局的日志中间件
	r.Use(gee.Logger())

	//静态文件
	//r.Static("/assets", "/usr/geektutu/blog/static")
	// 或相对路径 r.Static("/assets", "./static")

	//添加路由
	//r.GET("/index", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>index</h1>")
	//})

	//分组
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s,path:%s\n", c.Query("name"), c.Path)
		})
	}

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	//运行HTTP服务
	r.Run(":9999")

}
