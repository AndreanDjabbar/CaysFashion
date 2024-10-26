package validators

import (
	"strings"

	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/requests"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/repositories"
	"github.com/AndreanDjabbar/CaysFashion/backend/pkg/utils"
)

func ValidateLoginRequest(userLogin requests.LoginRequest) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(userLogin.Username) == "" {
		errors["username"] = "Username is required"
	} else if len(userLogin.Username) < 6 {
		errors["username"] = "Username must be at least 6 characters long"
	}

	if strings.TrimSpace(userLogin.Password) == "" {
		errors["password"] = "Password is required"
	} else if len(userLogin.Password) < 8 {
		errors["password"] = "Password must be at least 8 characters long"
	}

	if len(errors) == 0 {
		user, err := repositories.GetUserByUsername(userLogin.Username)
		if err != nil {
			errors["username"] = "User not found"
		}
		if valid := utils.CheckPasswordHash(userLogin.Password, user.Password); !valid {
			errors["password"] = "Incorrect password"
		}
	}
	return errors
}
