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
