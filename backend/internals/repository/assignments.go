package repository

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"time"
)

func AddNewAssignment(db *sql.DB, assignment *models.Assignment) error {
	query := `INSERT INTO assignments (name, class_id, statement, time_created, dead_line) VALUES(?, ?, ?, ?, ?)`

	timeCreated, err := time.Parse(time.DateTime, assignment.TimeCreated)
	deadLine, err := time.Parse(time.DateTime, assignment.DeadLine)

	_, err = db.Exec(
		query,
		assignment.Name,
		assignment.ClassId,
		assignment.Statement,
		timeCreated,
		deadLine)

	return err
}

func DeleteAssignment(db *sql.DB, assignmentId int) error {
	query := "DELETE FROM assignments WHERE id = ?"

	_, err := db.Exec(query, assignmentId)

	return err
}

func AddNewMark(db *sql.DB, mark models.Mark) error {
	query := "INSERT INTO marks (assignment_id, student_id, mark) VALUES(?, ?, ?)"

	_, err := db.Exec(query, mark.AssignmentId, mark.StudentId, mark.Mark)

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

func GetAssignmentByID(db *sql.DB, assignmentId int) (models.Assignment, error) {
	var assignment models.Assignment

	query := "SELECT id, name, class_id, statement, time_created, dead_line" +
		" FROM assignments WHERE id = ?"

	err := db.QueryRow(query, assignmentId).
		Scan(&assignment.Id,
			&assignment.Name,
			&assignment.ClassId,
			&assignment.Statement,
			&assignment.TimeCreated,
			&assignment.DeadLine)

	return assignment, err
}
