package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func ClassExists(db *sql.DB, classCode string) bool {
	var alreadyExists bool
	err := db.QueryRow("SElECT EXISTS (SELECT * FROM classes WHERE class_code = ?)", classCode).
		Scan(&alreadyExists)

	if err != nil {
		// TBD
		return false
	}

	return alreadyExists
}

func CreateNewClass(db *sql.DB, classCode, className string, teacherId int) error {
	query := `INSERT INTO classes (class_code, name, teacher_id) VALUES (?, ?, ?)`

	_, err := db.Exec(query, classCode, className, teacherId)

	return err
}

func GetClass(db *sql.DB, classCode string) (models.Class, error) {
	var class models.Class
	err := db.QueryRow("SELECT * FROM classes WHERE class_code = ?", classCode).
		Scan(&class.Id, &class.Code, &class.Name, &class.TeacherId)

	return class, err
}

func GetClassByCode(db *sql.DB, classCode string) (models.Class, error) {
	class, err := GetClass(db, classCode)

	if err != nil {
		return models.Class{}, err
	}

	return class, nil
}

func StudentExistInClass(db *sql.DB, studentId, classId int) bool {
	var alreadyExists bool
	err := db.QueryRow(
		"SElECT EXISTS (SELECT * FROM students_of_classes WHERE class_id = ? AND student_id = ?)",
		classId, studentId).Scan(&alreadyExists)

	if err != nil {
		// TBD
		return true
	}

	return alreadyExists
}

func JoinTheClass(db *sql.DB, studentId, classId int) error {
	if StudentExistInClass(db, studentId, classId) {
		return nil
	}

	query := `INSERT INTO students_of_classes (student_id, class_id) VALUES (?, ?)`

	_, err := db.Exec(query, studentId, classId)

	return err
}
