package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learnGo/learn_gin/controller"
	"learnGo/learn_gin/model"
)

var R = gin.Default()

func init() {
	//定义路由组
	user := R.Group("/user")
	{
		user.POST("/login", model.DoLogin)
		user.POST("/bindStruct", model.GetParamsByStruct)
		user.GET("/index", controller.UIndex)
	}

	g := R.Group("/g")
	{
		g.GET("/query", model.GetParamsByQuery)
		g.GET("/getQuery", model.GetParamsByGetQuery)
		g.GET("/defaultQuery", model.GetParamsByDefaultQuery)
	}

	p := R.Group("/p")
	{
		p.POST("/postForm", model.PostParamsByPostForm)
		p.POST("/getPostForm", model.PostParamsByGetPostForm)
		p.POST("/defaultPostForm", model.PostParamsByDefaultPostForm)
	}

	p2 := R.Group("/p2")
	{
		p2.GET("/:id", model.GetParamsByPath)
	}

	u := R.Group("/util")
	{
		//定义一个GET请求的路由，请求路径为：/ping
		u.GET("/ping", model.Pong)
		u.GET("/ip", model.ClientIp)
	}

	f := R.Group("/f")
	{
		f.GET("download", model.Download)
	}

	fmt.Println("----router init succeed...")
}
