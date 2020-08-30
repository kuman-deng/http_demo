package main

import (
	"fmt"
	"log"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is a test page"))
}

func getName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello kuman"))
}

func main() {
	fmt.Println("starting http server...")

	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/test", test)
	http.HandleFunc("/name", getName)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}