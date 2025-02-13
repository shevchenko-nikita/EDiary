package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"golang.org/x/crypto/bcrypt"
)

func AddNewUser(db *sql.DB, user *models.User) error {
	if err := ValidateUser(user); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("error occured while hashing password")
	}

	user.Password = string(hashedPassword)

	return repository.AddNewUser(db, user)
}
