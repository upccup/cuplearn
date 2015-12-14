package main

import (
	"net/http"
)

// 重定向服务器访问地址

func main() {
	http.ListenAndServe(":1434", http.RedirectHandler("http://www.baidu.com", 301))
}
