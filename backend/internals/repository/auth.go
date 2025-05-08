package repository

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func FindUserByUsername(db *sql.DB, username string) (*models.User, error) {
	var user models.User

	query := "SELECT id, first_name, middle_name, second_name, email, username, password, profile_image_path " +
		"FROM users WHERE username = ?"

	err := db.QueryRow(query, username).
		Scan(&user.ID, &user.FirstName, &user.MiddleName, &user.SecondName,
			&user.Email, &user.Username, &user.Password, &user.ProfileImgPath)

	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func AddNewUser(db *sql.DB, user *models.User) error {
	var alreadyExists bool

	err := db.QueryRow("SElECT EXISTS (SELECT * FROM users WHERE username = ?)", user.Username).Scan(&alreadyExists)

	if err != nil {
		return err
	}

	if alreadyExists {
		return fmt.Errorf("user already exists")
	}

	query := `INSERT INTO users (
                   first_name, 
                   middle_name, 
                   second_name, 
                   email, 
                   username, 
                   password, 
                   profile_image_path
              ) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)
			  `
	_, err = db.Exec(query,
		user.FirstName,
		user.MiddleName,
		user.SecondName,
		user.Email,
		user.Username,
		user.Password,
		user.ProfileImgPath,
	)

	if err != nil {
		return err
	}

	return nil
}
