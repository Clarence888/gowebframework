package main

import (
	"fmt"
	"log"
	"net/http"
)

//标准库启动一个web服务 因为go内置了  net/http库 封装了基本接口 所以可直接使用

func main() {

	//相当于注册路由到处理服务里 底层是用map实现 相当于把路由搞到map里
	//http.HandleFunc接受两个参数：第一个参数是字符串表示的 url 路径，第二个参数是该 url 实际的处理对象。
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/hello", helloHandler)

	//ListenAndServe第二个参数就是一个 Handler 函数
	//这里第二个参数传入 nil 可以想一下为什么
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {

	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
