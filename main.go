package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {

}
//使用自己的handler，接管所有的http请求，做统一处理
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w,"404 not found %s\n",req.URL)
	}
}

func main() {

	engine := new(Engine)

	log.Fatal(http.ListenAndServe(":9999", engine))
}
