package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type U struct {
	Name string `json:"name" form:"name"` //json表示转成json字符串时的字段名，form表示表单提交时的参数名
	Age  int    `json:"age" form:"age"`
}

func DoLogin(c *gin.Context) {
	uname := c.PostForm("username")
	paswd := c.PostForm("password")

	c.JSON(200, gin.H{
		"username": uname,
		"password": paswd,
	})
	//c.String(200, "uname=%s&pass=%s", uname, paswd)
}

func GetParamsByStruct(c *gin.Context) {
	var u U

	if c.ShouldBind(&u) == nil {
		c.JSON(http.StatusOK, gin.H{
			"name":   u.Name,
			"age":    u.Age,
			"status": "accepted",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"MSG": "NO PARAMS",
		})
	}
}
