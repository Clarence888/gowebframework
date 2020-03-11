package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URLPath = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9991")
}

//注意 在执行的时候 发现无法执行 是因为勾选了 vnedor模式 这样的话 就需要把依赖的包放到vendor目录下了

//这种方式不推荐了 所以 需要去掉勾选vendoring mode
