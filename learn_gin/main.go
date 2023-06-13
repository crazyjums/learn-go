package main

import (
	"github.com/gin-gonic/gin"
	"learnGo/learn_gin/router"
)

func main() {
	gin.SetMode(gin.DebugMode) //设置debug模式，会打印很多调试日志
	router.R.LoadHTMLGlob("/learn_gin/status/template/**/*")
	if err := router.R.Run(); err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
