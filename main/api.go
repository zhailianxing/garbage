package main

import (
	"garbage/Logic/garbage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

type returnData struct {
	name     string
	category string
}

func GetCategory(c *gin.Context) {
	//key := c.Request.FormValue("key")
	//key2, exist := c.Params.Get("key")
	//if exist {
	//	fmt.Println("key2:", key2)
	//}
	//fmt.Println("key:", key)

	key := c.PostForm("key") // 只支持  Content-Type: application/x-www-form-urlencoded
	data := garbage.GetCategoryByName(key)
	c.JSON(http.StatusOK, gin.H{
		"ret":  0,
		"msg":  "success",
		"list": data,
	})
}
