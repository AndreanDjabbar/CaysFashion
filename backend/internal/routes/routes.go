package routes

import (
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/auth/jwt"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpMainRoutes(router *gin.Engine) {
	router.GET("/", handlers.DefaultRootHandler)
	
	route := router.Group("/caysfashion")
	route.GET("/", handlers.MainRootHandler)
	route.POST("/register", handlers.RegisterHandler)
	route.POST("/login", handlers.LoginHandler)

	route.Use(jwt.AuthMiddleware()) 
	{
		route.GET("/home", handlers.HomeHandler)
	}
}