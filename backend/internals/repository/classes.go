package repository

import (
	"database/sql"
)

func ClassExists(db *sql.DB, classCode string) bool {
	var alreadyExists bool
	err := db.QueryRow("SElECT EXISTS (SELECT * FROM classes WHERE class_code_id = ?)", classCode).
		Scan(&alreadyExists)

	if err != nil {
		// TBD
		return false
	}

	return alreadyExists
}

func CreateNewClass(db *sql.DB, classCode, className string, teacherId int) error {
	query := `INSERT INTO classes (class_code_id, name, teacher_id) VALUES (?, ?, ?)`

	_, err := db.Exec(query, classCode, className, teacherId)

	return err
}
