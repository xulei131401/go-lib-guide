package main

import (
	"log"
	"net/http"
	"strings"
)

/*
*
结论；
 1. go内置了文件类型服务，可以很方便的启动一个本地文件服务
 2. 如果不想显示列表，可以目录里设置index.html
 3. http.StripPrefix()
 4. http.FileServer()
 5. http.Dir()
 6. 没有直接启动http.FileServer，而是通过路由分发模式启动
*/
func main() {
	fileServer := http.FileServer(http.Dir("static"))
	server := http.Server{
		Addr: ":8886",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("req.URL.String():", r.URL.String())
			if strings.HasPrefix(r.URL.String(), "/static") {
				http.StripPrefix("/static/", intercept(fileServer)).ServeHTTP(w, r)
				return
			}
			//fileServer.ServeHTTP(w, r)
			//http.StripPrefix("/", intercept(fileServer)).ServeHTTP(w, r)
			log.Println("非文件类型请求")
			return
		}),
	}

	log.Printf("服务启动 。。。")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("服务启动失败 %s \n", err)
	}
}

func intercept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("r.URL.String():%#v", *r.URL)
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
