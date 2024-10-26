package routes

import (
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpMainRoutes(router *gin.Engine) {
	router.GET("/", handlers.DefaultRootHandler)
	
	route := router.Group("/caysfashion")
	route.GET("/", handlers.MainRootHandler)
	route.POST("/register", handlers.RegisterHandler)
}

// func SetUpSpecialRoutes(router *gin.Engine) {
// 	store2 := cookie.NewStore([]byte("secret2"))
// 	store2.Options(sessions.Options{
// 		// different options
// 	})
// 	router.Use(sessions.Sessions("mysession", store2))
// 	// router.GET("/", ...)
// }