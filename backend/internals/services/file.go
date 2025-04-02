package services

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

func UploadFile(db *sql.DB, fileName, relativePath string, userID, assignmentID int) error {
	return repository.UploadFile(db, fileName, relativePath, userID, assignmentID)
}
