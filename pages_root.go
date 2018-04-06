package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": siteTitle,
	})
}

func rootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "page.html", gin.H{
		"title": siteTitle,
	})
}
