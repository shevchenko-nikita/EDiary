package services

import (
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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
	user.ProfileImgPath = os.Getenv("DEFAULT_IMAGE_PATH")

	return repository.AddNewUser(db, user)
}

func SignIn(db *sql.DB, username string, password string) (string, error) {
	user, err := repository.FindUserByUsername(db, username)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return GenerateToken(db, user)
}

func GenerateToken(db *sql.DB, user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", fmt.Errorf("error occured while signing token")
	}

	return tokenString, nil
}
