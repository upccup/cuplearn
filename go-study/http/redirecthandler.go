package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":1434", http.RedirectHandler("http://www.baidu.com", 301))
}
