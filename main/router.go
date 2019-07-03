package main

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/getCategory", GetCategory)

	//router.GET("/persons", GetPersonsApi)
	//
	//router.GET("/person/:id", GetPersonApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//
	//router.DELETE("/person/:id", DelPersonApi)

	return router
}
