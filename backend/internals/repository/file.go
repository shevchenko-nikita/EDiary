package repository

import "database/sql"

func UploadFile(db *sql.DB, fileName, relativePath string, userID, assignmentID int) error {
	query := "INSERT INTO students_files(file_name, file_path, assignment_id, student_id) VALUES(?, ?, ?, ?)"
	_, err := db.Exec(query, fileName, relativePath, assignmentID, userID)
	return err
}
