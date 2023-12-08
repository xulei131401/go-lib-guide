package main

import (
	"io/fs"
	"log"
	"net/http"
	"strings"
)

/*
*
embed FileServer结论；
 1. /template 允许访问 会重定向 2次
 2. /template/ 不会重定向
 3. /template/index.html 允许访问 会重定向 2次
 4. 首页还会显示目录
 5. 需要配合fs.Sub()返回下层目录使用
*/
func main() {
	aa, _ := fs.Sub(TemplateFs, "template")
	fileServer := http.FileServer(http.FS(aa))
	server := http.Server{
		Addr: ":8886",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("req.URL.String():", r.URL.String())
			//if strings.HasPrefix(r.URL.String(), "/static") {
			//	http.StripPrefix("/static/", intercept(fileServer)).ServeHTTP(w, r)
			//	return
			//}
			fileServer.ServeHTTP(w, r)
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
