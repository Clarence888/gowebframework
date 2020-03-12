package main

//设计上下文 web服务 请求 响应 都要处理 而且尽量不要让用户每次都去设置相关的信息
//比如返回json 返回html

//封装前返回json
/*
obj = map[string]interface{}{
"name": "geektutu",
"password": "1234",
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
encoder := json.NewEncoder(w)
...

*/

//封装后
/*
c.JSON(http.StatusOk,gee.H{
	"aaa":c.PostForm("aaa")
})
*/
