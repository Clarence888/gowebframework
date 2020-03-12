package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	/*
		r.GET("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "URLPath = %q\n", req.URL.Path)
		})
		r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
			for k, v := range req.Header {
				fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			}
		})
	*/

	//优化

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1> hello world</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9991")
}

//注意 在执行的时候 发现无法执行 是因为勾选了 vnedor模式 这样的话 就需要把依赖的包放到vendor目录下了

//这种方式不推荐了 所以 需要去掉勾选vendoring mode

/*
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/
