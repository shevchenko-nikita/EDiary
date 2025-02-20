package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
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
