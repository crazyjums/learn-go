package model

import "github.com/gin-gonic/gin"

func Download(c *gin.Context) {
	//直接返回一个文件，可以用来做文件下载。
	//c.FileAttachment("文件路径", "下载的文件名")
	//过File函数，直接返回本地文件，参数为本地文件地址。
	//c.File("C:\\Users\\zhuhonggen-hj\\Desktop\\20154241853060994.jpg")

	//设置请求头
	c.Header("site", "1234")
	c.FileAttachment("C:\\Users\\zhuhonggen-hj\\Desktop\\20154241853060994.jpg",
		"123.jpg")
}
