package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "template/user/index", gin.H{
		"title": "用户模板title",
	})
}
