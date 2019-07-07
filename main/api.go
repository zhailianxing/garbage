package main

import (
	"garbage/Logic/garbage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func GetCategory(c *gin.Context) {
	//key := c.Request.FormValue("key")
	//key2, exist := c.Params.Get("key")
	//if exist {
	//	fmt.Println("key2:", key2)
	//}
	//fmt.Println("key:", key)

	key := c.PostForm("key") // 只支持  Content-Type: application/x-www-form-urlencoded
	if len(key) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"ret":  -1,
			"msg":  "error",
			"list": make([]*garbage.ReturnData, 0),
		})
	} else {
		data := garbage.GetCategoryByName(key)
		isFound := false
		if len(data) > 0 {
			isFound = true
		}
		garbage.RecordSearchWords(key, isFound)
		c.JSON(http.StatusOK, gin.H{
			"ret":  0,
			"msg":  "success",
			"list": data,
		})
	}

}
