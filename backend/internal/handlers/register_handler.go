package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AndreanDjabbar/CaysFashion/backend/database/redis_service"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/entities"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/requests"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/repositories"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/validators"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	client = redis_service.GetClient()
	ctx    = redis_service.GetRedisContext()
	redisKey = "register-redis-key"
)

func RegisterHandler(c *gin.Context) {
	var registerRequest requests.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format", "details": err.Error()})
		return
	}

	if validationErrors := validators.ValidateRegisterRequest(registerRequest); len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation errors",
			"errors":  validationErrors,
		})
		return
	}

	otpCode, err := utils.GenerateOTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP code"})
		return
	}

	subject := "CaysFashion Email Verification"
	body := fmt.Sprintf(`
		<html>
		<body>
			<p>Your OTP Code is: <strong>%s</strong></p>
			<p>Valid for 5 minutes. If you did not request this, ignore this email.</p>
		</body>
		</html>
	`, otpCode)

	if err := utils.SendEmail(registerRequest.Email, "smtp.gmail.com", body, subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP code"})
		return
	}

	hashedPassword, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	randomToken := utils.GenerateRandomElements(50)

	err = client.HSet(ctx, redisKey, map[string]interface{}{
		"username": registerRequest.Username,
		"email":    registerRequest.Email,
		"password": hashedPassword,
		"otp":      otpCode,
		"randomToken": randomToken,
	}).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store user data"})
		return
	}

	if err := client.Expire(ctx, redisKey, 5*time.Minute).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set expiration"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"data": gin.H{
			"randomToken": randomToken,
		},
		"message": "Register form inputted successfully",
	})
}

func VerifyOTPHandler(c *gin.Context) {
	randomToken := c.Query("randomToken")
	if randomToken == "" {
		log.Error("Random token is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Random token is required"})
		return 
	}

	dataRedis, err := client.HGetAll(ctx, redisKey).Result() 

	randomTokenRedis, ok := dataRedis["randomToken"]
	if !ok {
		log.Error("Random token does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid random token"})
		return
	}

	if err == redis.Nil {
		log.Error("Random token does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid random token"})
		return
	} else if err != nil {
		log.Error("Failed to get random token from Redis", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get random token"})
		return
	}

	if randomToken != randomTokenRedis {
		log.Error("Invalid random token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid random token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Random token verified successfully"})
}

func OTPVerificationHandler(c *gin.Context) {
	var otpReq requests.OTPRequest

	if err := c.ShouldBindJSON(&otpReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP format"})
		return
	}

	dataRedis, err := client.HGetAll(ctx, redisKey).Result()
	if err == redis.Nil {
		log.Error("Random token does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid random token"})
		return
	} else if err != nil {
		log.Error("Failed to get random token from Redis", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get random token"})
	} 

	expectedOTP, ok := dataRedis["otp"]
	if !ok {
		log.Error("Random token does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid random token"})
		return
	}

	if otpReq.OTP == expectedOTP {
		RegisterUser(c, dataRedis)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
	}
}

func RegisterUser(c *gin.Context, dataRedis map[string]string) {
	var newUser entities.User
	username, ok := dataRedis["username"]
	if !ok {
		log.Error("Username does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	email, ok := dataRedis["email"]
	if !ok {
		log.Error("Email does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	password, ok := dataRedis["password"]
	if !ok {
		log.Error("Password does not exist in Redis")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	newUser = entities.User{
		Username: username,
		Email:    email,
		Password: password,
		Role: "User",
	}

	if err := repositories.CreateUser(&newUser); err != nil {
		log.Error("Failed to create user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}