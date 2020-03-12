package gee

//框架入口

import (
	"net/http"
)

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type HandlerFunc func(c *Context)

//路由映射表 用来管理路由
type Engine struct {
	//router map[string]HandlerFunc
	router *router
}

//相当于gee engine的构造方法
func New() *Engine {
	//return &Engine{router: make(map[string]HandlerFunc)}
	return &Engine{router: newRouter()}
}

//添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//为什么这么设计呢 因为 相同的路由 不同请求方法 可以映射不同的处理方法 但是尽量不要这么去做
	//key := method + "-" + pattern
	//engine.router[key] = handler

	//优化
	engine.router.addRoute(method, pattern, handler)
}

//添加get请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

//添加post请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//实现接口定义的ServeHTTP方法 engine 才能传给http.ListenAndServe 第二个参数
//解析路径 找到映射表 找到相关的处理方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	/*
		key := req.Method + "-" + req.URL.Path
		if handler, ok := engine.router[key]; ok {
			handler(w, req)
		} else {
			fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
		}
	*/
	//优化
	c := newContext(w, req)
	engine.router.handle(c)
}

//添加启动方法

func (engine *Engine) Run(addr string) (err error) {
	//注意 如果没有实现ServeHTTP方法 此处第二个参数会报错
	return http.ListenAndServe(addr, engine)
}
