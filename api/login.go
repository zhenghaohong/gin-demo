package api

import "github.com/gin-gonic/gin"

func V1login(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "v1 login",
	})
	return
}

func V2login(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "v2 login",
	})
	return
}

