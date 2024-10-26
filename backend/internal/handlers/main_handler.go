package handlers

import (
	"net/http"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.SetUpLogger()

func DefaultRootHandler(c *gin.Context) {
	c.Redirect(
		http.StatusMovedPermanently,
		"/caysfashion",
	)
}

func MainRootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to CaysFashion",
	})
}

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the home page",
	})
}