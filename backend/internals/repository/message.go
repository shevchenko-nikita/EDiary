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

func DeleteClassMessage(db *sql.DB, messageId int) error {
	query := "DELETE FROM class_comments WHERE id = ?"

	_, err := db.Exec(query, messageId)

	return err
}

func GetAllClassMessages(db *sql.DB, classId int) ([]models.Message, error) {
	query := "SELECT id, class_id, user_id, text, time_posted FROM class_comments WHERE class_id = ?"

	rows, err := db.Query(query, classId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message

		err := rows.Scan(
			&message.Id,
			&message.ClassId,
			&message.UserId,
			&message.Text,
			&message.TimePosted)
		
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func MessageExists(db *sql.DB, messageId int) (bool, error) {
	query := "SELECT EXISTS(SELECT * FROM class_comments WHERE id = ?)"

	var exists bool

	err := db.QueryRow(query, messageId).Scan(&exists)

	return exists, err
}

func GetMessageById(db *sql.DB, messageId int) (models.Message, error) {
	var message models.Message
	query := "SELECT * FROM class_comments WHERE id = ?"

	row := db.QueryRow(query, messageId)

	err := row.Scan(&message.Id, &message.ClassId, &message.UserId, &message.Text, &message.TimePosted)

	return message, err
}
