package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"os"
)

func GetUserById(db *sql.DB, userId int) (models.User, error) {
	query := "SELECT id, first_name, middle_name, second_name, email, username, password, profile_image_path " +
		"FROM users WHERE id = ?"

	var user models.User

	err := db.QueryRow(query, userId).
		Scan(
			&user.Id,
			&user.FirstName,
			&user.MiddleName,
			&user.SecondName,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.ProfileImgPath)

	return user, err
}

func UpdateUserProfile(db *sql.DB, newUserInfo *models.User) error {
	query := "UPDATE users " +
		"SET first_name = ?, middle_name = ?, second_name = ?, email = ?, username = ? WHERE id = ?"

	_, err := db.Exec(
		query,
		newUserInfo.FirstName,
		newUserInfo.MiddleName,
		newUserInfo.SecondName,
		newUserInfo.Email,
		newUserInfo.Username,
		newUserInfo.Id)

	return err
}

func UpdateUserProfileImage(db *sql.DB, userId int, imageDst string) error {
	query := "UPDATE users SET profile_image_path = ? WHERE id = ?"

	_, err := db.Exec(query, imageDst, userId)

	return err
}

func DeleteUserProfileImage(db *sql.DB, userId int) error {
	query := "UPDATE users SET profile_image_path = ? WHERE id = ?"

	defaultImgPath := os.Getenv("DEFAULT_IMAGE_PATH")

	_, err := db.Exec(query, defaultImgPath, userId)

	return err
}

func UserExists(db *sql.DB, userId int) (bool, error) {
	query := "SELECT EXISTS(SELECT * FROM users WHERE id = ?)"

	var exists bool

	err := db.QueryRow(query, userId).Scan(&exists)

	return exists, err
}
