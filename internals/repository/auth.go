package repository

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func AddNewUser(db *sql.DB, user *models.User) error {
	var alreadyExists bool

	err := db.QueryRow("SElECT EXISTS (SELECT * FROM users WHERE username = ?)", user.Username).Scan(&alreadyExists)

	if err != nil {
		return fmt.Errorf("error occured while insert user to DB: %v", err)
	}

	if alreadyExists {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	query := `INSERT INTO users (
                   first_name, 
                   father_name, 
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
		user.FatherName,
		user.SecondName,
		user.Email,
		user.Username,
		user.Password,
		user.ProfileImgPath,
	)

	if err != nil {
		return fmt.Errorf("failed to insert user to DB: %v", err)
	}

	return nil
}
