package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func AddNewMark(db *sql.DB, mark models.Mark) error {
	query := "INSERT INTO marks (class_id, assignment_id, student_id, mark) VALUES(?, ?, ?, ?)"

	_, err := db.Exec(query, mark.ClassID, mark.AssignmentID, mark.StudentID, mark.Mark)

	return err
}

func UpdateMark(db *sql.DB, mark models.Mark) error {
	query := "UPDATE marks SET mark = ? WHERE assignment_id = ? AND student_id = ?"

	_, err := db.Exec(query, mark.Mark, mark.AssignmentID, mark.StudentID)

	return err
}

func MarkAlreadyExist(db *sql.DB, mark models.Mark) bool {
	query := "SElECT EXISTS (SELECT * FROM marks WHERE assignment_id = ? AND student_id = ?)"

	var alreadyExist bool

	if err := db.QueryRow(query, mark.AssignmentID, mark.StudentID).Scan(&alreadyExist); err != nil {
		return false
	}

	return alreadyExist
}

func GetMark(db *sql.DB, userID, assignmentID int) (int, error) {
	query := "SELECT mark FROM marks WHERE assignment_id = ? AND student_id = ?"
	var mark int
	if err := db.QueryRow(query, assignmentID, userID).Scan(&mark); err != nil {
		return 0, err
	}

	return mark, nil
}
