package services

import (
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func ValidatePassword(password string) error {
	if len(password) < 5 || len(password) > 40 {
		return fmt.Errorf("password must be between 8 and 40 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	if len(email) == 0 {
		return fmt.Errorf("empty email")
	}
	return nil
}

func ValidateUsername(username string) error {
	if len(username) < 5 || len(username) > 20 {
		return fmt.Errorf("username must be between 5 and 20 characters")
	}
	return nil
}

func ValidateUser(user *models.User) error {
	if err := ValidateEmail(user.Email); err != nil {
		return fmt.Errorf("Invalid email address")
	}

	if err := ValidateUsername(user.Username); err != nil {
		return fmt.Errorf("Invalid username")
	}

	if err := ValidatePassword(user.Password); err != nil {
		return fmt.Errorf("Invalid password")
	}

	return nil
}
