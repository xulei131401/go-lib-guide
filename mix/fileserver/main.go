package main

import (
	"log"
	"strings"
)

func main() {
	// 测试路径函数
	log.Println(strings.HasSuffix("/static/", "/"))
	log.Println(strings.HasSuffix("/", "/"))

}
