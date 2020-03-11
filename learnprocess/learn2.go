package main

import (
	"fmt"
	"log"
	"net/http"
)

//创建一个结构体 可以处理所有的路由
type Engine struct{}

//用该结构体实现了方法
func (engine Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	//拦截了所有的HTTP请求，拥有了统一的控制入口
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		//fmt.Println(req.Header) 会发现输出的header信息每次顺序都不同 是因为 req.Header是个map
		//map为啥 用range循环输出的时候不是固定的呢？这里有相关的底层知识。 主要是因为map的底层实现决定的
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]= %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 no found: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9998", engine))
}
