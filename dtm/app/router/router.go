package router

import (
	"github.com/gin-gonic/gin"
	"holygin/app/controller/student"
	"holygin/app/middleware"
)

func InitRouter(router *gin.Engine) {
	registerMiddleware(router)
	registerRouter(router)
}

func registerMiddleware(router *gin.Engine) {
	router.Use(
		middleware.ParameterMiddleware(),
		middleware.DevMiddleware(),
	)
}

func registerRouter(router *gin.Engine) {
	router.Static("/assets", "../static/assets")
	router.LoadHTMLGlob("app/templates/*")
	/*var files []string
	_ = filepath.Walk("../templates/", func(path string, info os.FileInfo, err error) error {
		fmt.Println("path:",path)
		if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".tmpl") {
			files = append(files, path)
		}
		return nil
	})*/

	//router.LoadHTMLFiles(files...)

	v1 := router.Group("/api")
	{
		v1.GET("/student/detail", student.DetailAction)
		v1.GET("/student/detail/tpl", student.DetailTplAction)
		v1.POST("/student/update", student.UpdateAction)
	}

	AddDTMRouter(router)
}