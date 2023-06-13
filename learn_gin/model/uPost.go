package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostParamsByPostForm(c *gin.Context) {
	name := c.PostForm("name")

	name2 := c.Query("name1")

	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"name2": name2,
	})
}

func PostParamsByGetPostForm(c *gin.Context) {
	name, ok := c.GetPostForm("name")

	if ok {
		c.JSON(http.StatusOK, gin.H{
			"name": "name is not accepted by post",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	}
}

func PostParamsByDefaultPostForm(c *gin.Context) {
	name := c.DefaultPostForm("name", "post默认值")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
