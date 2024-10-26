package handlers

import (
	"net/http"

	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/entities"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/requests"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/repositories"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/validators"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var registerRequest requests.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		log.Error("Invalid input format", "error", err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	if validationErrors := validators.ValidateRegisterRequest(registerRequest); len(validationErrors) > 0 {
		log.Error("Validation errors", "errors", validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation errors",
			"errors":  validationErrors,
		})
		return
	}

	log.Info("Registering user", "username", registerRequest.Username)

	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		log.Error("Could not hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	user := entities.User{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Password: hashedPassword,
		Role:     entities.UserRoleUser,
	}

	if err := repositories.CreateUser(&user); err != nil {
		log.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	log.Info("User registered successfully", "userID", user.UserID)
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"userID": user.UserID,
		},
		"message": "User registered successfully",
	})
}
