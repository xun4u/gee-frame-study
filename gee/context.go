package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//定义别名H，用于构建json数据更方便
type H map[string]interface{}

//定义上下文的构造体：包含了请求和响应等相关信息
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

//实例化
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

//获取get参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//获取表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//写入响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//设置响应头
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

//文本响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//json响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//字节流
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

//html响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
