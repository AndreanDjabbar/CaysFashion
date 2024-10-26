package jwt

import (
	"net/http"
	"strings"

	"github.com/AndreanDjabbar/CaysFashion/backend/internal/middlewares"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.SetUpLogger()

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
			log.Error("Missing or invalid auth token")
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid auth token"})
            c.Abort()
            return
        }

        if strings.HasPrefix(tokenString, "Bearer ") {
            tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid auth token"})
            c.Abort() 
            return
        }

        claims, err := middlewares.ValidateToken(tokenString)
        if err != nil {
            log.Error("Invalid token", "error", err)
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
            c.Abort() 
            return
        }

        c.Set("claims", claims) 
        c.Next()
    }
}