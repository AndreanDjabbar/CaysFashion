package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRootHandler(c *gin.Context) {
	c.Redirect(
		http.StatusMovedPermanently,
		"/caysfashion",
	)
}

func MainRootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}