package validators

import (
	"regexp"
	"strings"

	"github.com/AndreanDjabbar/CaysFashion/backend/internal/models/requests"
	"github.com/AndreanDjabbar/CaysFashion/backend/internal/repositories"
)

func ValidateUserRegister(userRegister requests.UserRegister) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(userRegister.Username) == "" {
		errors["username"] = "Username is required"
	}

	if (strings.TrimSpace(userRegister.Username) != "") && (len(userRegister.Username) < 6) {
		errors["username"] = "Username must be at least 6 characters long"
	}

	if strings.TrimSpace(userRegister.Email) == "" {
		errors["email"] = "Email is required"
	}

	if (strings.TrimSpace(userRegister.Email) == "") && (!isValidEmail(userRegister.Email)) {
		errors["email"] = "Email must be a valid email address"
	}

	if strings.TrimSpace(userRegister.Password) == "" {
		errors["password"] = "password is required"
	}

	if (strings.TrimSpace(userRegister.Password) == "") && (len(userRegister.Password) < 8) {
		errors["password"] = "Password must be at least 8 characters long"
	}

	if _, err := repositories.GetUserByUsername(userRegister.Username); err == nil {
		errors["username"] = "Username is already taken"
	}

	if _, err := repositories.GetUserByEmail(userRegister.Email); err == nil {
		errors["email"] = "Email is already taken"
	}

	return errors
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
