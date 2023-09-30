package main

import (
	"log"
	"net/http"
)

/*
*
结论：
 1. 直接启动http.FileServer
*/
func main() {
	log.Printf("服务启动 127.0.0.1:8886 。。。")
	http.Handle("/", http.FileServer(http.Dir("static")))
	if err := http.ListenAndServe(":8886", nil); err != nil {
		log.Fatalf("服务启动失败 %s \n", err)
	}
}
