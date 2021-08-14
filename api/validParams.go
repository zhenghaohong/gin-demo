package api

import (
	"encoding/json"
	"fmt"
	"gin-demo1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// 验证参数格式


type Part struct {
	Id int `form:"id" json:"id"`
	Name string `json:"name" binding:"required"`
	Num  int `json:"num" binding:"required"`
}


func CreatePart(c *gin.Context) {
	var params Part
	if err := c.ShouldBind(&params); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			errFieldType := reflect.TypeOf(err)
			errErrorType := reflect.TypeOf(err.Error())

			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
				"message":"参数类型有误",
				"errFieldType":errFieldType,
				"errErrorType":errErrorType,
				"err":err,
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(utils.Trans),
		})
		return
	}

}


type ParamsValidError struct {
	Value string
	Type  interface{}
	Offset int
	Struct string
	Field  string
}


// map 转成 struct
func MapToStruct(data string) {

	var m map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatal(err)
	}

	var p ParamsValidError
	mapstructure.Decode(m, &p)
	fmt.Println("ParamsValidError", p)

}



// 1 interface to Struct
func InterfaceToStruct1(data interface{}) {
	// newData := data
	// var newInterface1 interface{}
	//
	// 第一种使用interface
	// newInterface1=newData
	// fmt.Printf("使用interface: %v",newInterface1.(ParamsValidError))

	p, ok := (data).(ParamsValidError)
	if ok {
		fmt.Println("id:" + p.Field)
		fmt.Println("name:" + p.Struct)
	} else {
		fmt.Println("can not convert")
	}

}


