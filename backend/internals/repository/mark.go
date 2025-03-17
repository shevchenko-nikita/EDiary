package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
)

func AddNewMark(db *sql.DB, mark models.Mark) error {
	query := "INSERT INTO marks (class_id, assignment_id, student_id, mark) VALUES(?, ?, ?, ?)"

	_, err := db.Exec(query, mark.ClassId, mark.AssignmentId, mark.StudentId, mark.Mark)

	return err
}

func UpdateMark(db *sql.DB, mark models.Mark) error {
	query := "UPDATE marks SET mark = ? WHERE assignment_id = ? AND student_id = ?"

	_, err := db.Exec(query, mark.Mark, mark.AssignmentId, mark.StudentId)

	return err
}

func MarkAlreadyExist(db *sql.DB, mark models.Mark) bool {
	query := "SElECT EXISTS (SELECT * FROM marks WHERE assignment_id = ? AND student_id = ?)"

	var alreadyExist bool

	if err := db.QueryRow(query, mark.AssignmentId, mark.StudentId).Scan(&alreadyExist); err != nil {
		return false
	}

	return alreadyExist
}
