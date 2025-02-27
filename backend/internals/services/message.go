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
