package repositories

import (
	"github.com/AndreanDjabbar/CaysFashion/backend/database/mysql_service"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/entities"
)

var db = mysql_service.GetDB()

func CreateUser(user *entities.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}