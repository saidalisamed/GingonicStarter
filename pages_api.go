package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
