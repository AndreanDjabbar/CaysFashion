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
	var userRegister requests.UserRegister
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		log.Error("Invalid input format", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	validationErrors := validators.ValidateUserRegister(userRegister)
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "errors": validationErrors})
		return
	}

	log.Info("Registering user", "username", userRegister.Username)
	hashedPassword, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		log.Error("Could not hash password", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	user := entities.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: hashedPassword,
		Role:     entities.UserRoleUser,
	}

	if err := repositories.CreateUser(&user); err != nil {
		log.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user"})
		return
	}

	log.Info("User registered successfully", "userID", user.UserID)
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"data": gin.H{
			"userID": user.UserID,
		},
		"message": "User registered successfully",
	})
}
