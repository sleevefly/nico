package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//加入服务器解析不了这儿的 ip地址需要和请求时的ip地址一样
	r.Run("0.0.0.0:9002")
}
