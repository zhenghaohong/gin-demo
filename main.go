package main

import (
"fmt"
	"gin-demo1/api"
	"github.com/gin-gonic/gin"

)

func main(){
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.GET("/v1login", api.V1login)
	r.GET("/v2login", api.V2login)


	r.POST("/part",api.CreatePart)
	r.POST("/importExcel",api.ReadExcel)

	err := r.Run(":8000")
	if err != nil {
		fmt.Println("server start failed ！！！")
	}

}




