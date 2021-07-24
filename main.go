package main

import (
"fmt"
"github.com/gin-gonic/gin"

)

func main(){
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.GET("/v1login", v1login)
	r.GET("/v2login", v2login)


	err := r.Run(":8000")
	if err != nil {
		fmt.Println("server start failed ！！！")
	}

}



func v1login(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "v1 login",
	})
	return
}

func v2login(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "v2 login",
	})
	return
}


