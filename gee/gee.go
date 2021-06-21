package gee

import (
	"fmt"
	"net/http"
)

//定义路由的方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	//路由表  地址=>handler
	router map[string]HandlerFunc
}

//构造函数
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

//增加路由条目到路由表
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

//添加get路由
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

//添加post路由
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

//开启服务
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

//使用自己的handler，接管所有的http请求，做统一处理
//通过路由表映射关系来找路由
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 not found %s\n", req.URL)
	}
}
