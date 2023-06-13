package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ClientIp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ip": c.ClientIP(),
	})
}
