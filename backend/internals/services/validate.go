package services

import (
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func ValidatePassword(password string) error {

	return nil
}

func ValidateEmail(email string) error {

	return nil
}

func ValidateUsername(username string) error {
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
