package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"time"
)

func CreateClassMessage(db *sql.DB, message models.Message) error {
	query := "INSERT INTO class_comments (class_id, user_id, text, time_posted) VALUES (?, ?, ?, ?)"

	timeCreated, err := time.Parse(time.DateTime, message.TimePosted)

	if err != nil {
		return err
	}

	_, err = db.Exec(query, message.ClassId, message.UserId, message.Text, timeCreated)

	return err
}
