package gee

import (
	"net/http"
)

//定义路由的方法
type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

//构造函数
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

//增加路由条目到路由表
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
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
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}
