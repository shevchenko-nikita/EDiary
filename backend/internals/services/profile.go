package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

func UpdateUserProfile(db *sql.DB, userId int, newUserInfo *models.User) error {
	if userId != newUserInfo.Id {
		return fmt.Errorf("User ID does not match")
	}

	exists, err := repository.UserExists(db, userId)

	if err != nil || !exists {
		return fmt.Errorf("User does not exist")
	}

	return repository.UpdateUserProfile(db, newUserInfo)
}
