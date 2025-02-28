package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

func CreateClassMessage(db *sql.DB, message models.Message) error {
	class, err := repository.GetClassById(db, message.ClassId)

	if err != nil {
		return err
	}

	if !repository.StudentExistInClass(db, message.UserId, class.Id) &&
		message.UserId != class.TeacherId {
		return fmt.Errorf("user doesn't have access")
	}

	return repository.CreateClassMessage(db, message)
}

func DeleteClassMessage(db *sql.DB, userId, messageId int) error {
	exist, err := repository.MessageExists(db, messageId)

	if !exist || err != nil {
		return fmt.Errorf("message doesn't exist")
	}

	message, err := repository.GetMessageById(db, messageId)

	if err != nil {
		return err
	}

	if message.UserId != userId {
		return fmt.Errorf("user doesn't have access")
	}

	return repository.DeleteClassMessage(db, messageId)
}

func GetAllClassMessages(db *sql.DB, userId, classId int) ([]models.Message, error) {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return nil, err
	}

	exists := repository.StudentExistInClass(db, userId, class.Id)

	if !exists && class.TeacherId != userId {
		return nil, fmt.Errorf("user doesn't have access")
	}

	return repository.GetAllClassMessages(db, classId)
}
