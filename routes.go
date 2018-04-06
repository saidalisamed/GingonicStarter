package main

import "github.com/gin-gonic/gin"

func routesConfig(router *gin.Engine) {

	// Root pages
	root := router.Group("/")
	{
		root.GET("/", index)
		root.GET("/page/*path", rootPage)
	}

	// API pages
	api := router.Group("/api")
	{
		api.POST("/email", apiEmail)
	}
}
