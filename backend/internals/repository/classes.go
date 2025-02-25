package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func CreateNewClass(db *sql.DB, classCode, className string, teacherId int) error {
	query := `INSERT INTO classes (class_code, name, teacher_id) VALUES (?, ?, ?)`

	_, err := db.Exec(query, classCode, className, teacherId)

	return err
}

func JoinTheClass(db *sql.DB, studentId, classId int) error {
	if StudentExistInClass(db, studentId, classId) {
		return nil
	}

	query := `INSERT INTO students_of_classes (student_id, class_id) VALUES (?, ?)`

	_, err := db.Exec(query, studentId, classId)

	return err
}

func UpdateClass(db *sql.DB, classId int, newClassName string) error {
	query := "UPDATE classes SET name = ? WHERE id = ?"

	_, err := db.Exec(query, newClassName, classId)

	return err
}

func DeleteClass(db *sql.DB, classId int) error {
	query := `DELETE FROM classes WHERE id = ?`

	_, err := db.Exec(query, classId)

	return err
}

func LeaveClass(db *sql.DB, studentId, classId int) error {
	query := `DELETE FROM students_of_classes WHERE student_id = ? AND class_id = ?`

	_, err := db.Exec(query, studentId, classId)

	return err
}

func GetClassById(db *sql.DB, classId int) (models.Class, error) {
	var class models.Class

	query := "SELECT id, class_code, name, teacher_id FROM classes WHERE id = ?"

	err := db.QueryRow(query, classId).
		Scan(&class.Id, &class.Code, &class.Name, &class.TeacherId)

	return class, err
}

func GetClassByCode(db *sql.DB, classCode string) (models.Class, error) {
	var class models.Class

	query := "SELECT id, class_code, name, teacher_id FROM classes WHERE class_code = ?"

	err := db.QueryRow(query, classCode).
		Scan(&class.Id, &class.Code, &class.Name, &class.TeacherId)

	return class, err
}

func GetStudentsList(db *sql.DB, classId int) ([]models.User, error) {
	query := "SELECT student_id FROM students_of_classes WHERE class_id = ?"

	rows, err := db.Query(query, classId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []models.User

	for rows.Next() {

		var userId int
		if err := rows.Scan(&userId); err != nil {
			return nil, err
		}

		user, err := GetUserById(db, userId)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
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
