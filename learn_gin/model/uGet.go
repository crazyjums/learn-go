package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetParamsByQuery(c *gin.Context) {
	name := c.Query("name") //type of name is string

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func GetParamsByGetQuery(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "name is not accepted",
		})
	}
}

func GetParamsByDefaultQuery(c *gin.Context) {
	name := c.DefaultQuery("name", "默认值")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
