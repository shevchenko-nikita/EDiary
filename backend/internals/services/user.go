package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"os"
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

func DeleteProfileImage(db *sql.DB, userId int) error {
	user, err := repository.GetUserById(db, userId)
	if err != nil {
		return err
	}

	if user.ProfileImgPath == os.Getenv("DEFAULT_IMAGE_PATH") {
		return nil
	}

	return repository.DeleteUserProfileImage(db, user.Id)
}

func UpdateUserProfileImage(db *sql.DB, userId int, imageDst string) error {
	return repository.UpdateUserProfileImage(db, userId, imageDst)
}
