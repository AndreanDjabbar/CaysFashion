package handlers

import (
	"net/http"

	"github.com/AndreanDjabbar/CaysFashion/backend/internal/middlewares"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/requests"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/repositories"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/validators"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		log.Error(
			"Invalid input format",
			"error", err,
		)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	if validationErrors := validators.ValidateLoginRequest(loginRequest); len(validationErrors) > 0 {
		log.Error("Validation errors", "errors", validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation errors",
			"errors":  validationErrors,
		})
		return
	}

	user, err := repositories.GetUserByUsername(loginRequest.Username)
	if err != nil {
		log.Error(
			"Failed to get user by username",
			"error", err,
		)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": gin.H{"message": "Failed to get user"}})
		return
	}

	token, err := middlewares.GenerateJWT(user.Username)
	if err != nil {
		log.Error(
			"Failed to generate token",
			"error", err,
		)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": gin.H{"message": "Failed to generate token"}})
		return
	}

	log.Info(
		"User logged in",
		"username", user.Username,
	)
	
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"userID":   user.UserID,
			"username": user.Username,
			"email":    user.Email,
			"token":    token,
		},
		"message": "Login successful",
	})
}
