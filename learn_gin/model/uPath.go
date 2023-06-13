package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetParamsByPath(c * gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
