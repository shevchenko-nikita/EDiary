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
		assignment.ClassID,
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
		Scan(&assignment.ID,
			&assignment.Name,
			&assignment.ClassID,
			&assignment.Statement,
			&assignment.TimeCreated,
			&assignment.DeadLine)

	return assignment, err
}

func UpdateAssignment(db *sql.DB, newAssignmentInfo *models.Assignment) error {
	query := "UPDATE assignments SET name = ?, statement = ?, dead_line = ? WHERE id = ?"

	deadLine, err := time.Parse(time.DateTime, newAssignmentInfo.DeadLine)

	if err != nil {
		return err
	}

	_, err = db.Exec(
		query,
		newAssignmentInfo.Name,
		newAssignmentInfo.Statement,
		deadLine,
		newAssignmentInfo.ID)

	return err
}

func GetAssignmentsList(db *sql.DB, classId int) ([]models.Assignment, error) {
	var assignments []models.Assignment

	query := "SELECT id, name, class_id, statement, time_created, dead_line FROM assignments WHERE class_id = ?"

	rows, err := db.Query(query, classId)

	if err != nil {
		return assignments, err
	}

	defer rows.Close()

	for rows.Next() {
		var assignment models.Assignment

		err := rows.Scan(
			&assignment.ID,
			&assignment.Name,
			&assignment.ClassID,
			&assignment.Statement,
			&assignment.TimeCreated,
			&assignment.DeadLine)

		if err != nil {
			return nil, err
		}

		assignments = append(assignments, assignment)
	}

	return assignments, nil
}

func AssignmentExist(db *sql.DB, assignmentId int) (bool, error) {
	var exist bool

	query := "SELECT EXISTS (SELECT * FROM assignments WHERE id = ?)"

	err := db.QueryRow(query, assignmentId).Scan(&exist)

	return exist, err
}

func GetAllClassMarks(db *sql.DB, classId int) ([]models.Mark, error) {
	var marks []models.Mark
	query := "SELECT m.id, m.class_id, m.assignment_id, m.student_id, m.mark " +
		"FROM marks m " +
		"JOIN assignments a ON m.assignment_id = a.id " +
		"WHERE a.class_id = ?"

	rows, err := db.Query(query, classId)
	if err != nil {
		return marks, err
	}

	defer rows.Close()

	for rows.Next() {
		var mark models.Mark

		err := rows.Scan(
			&mark.ID,
			&mark.ClassID,
			&mark.AssignmentID,
			&mark.StudentID,
			&mark.Mark)

		if err != nil {
			return marks, err
		}

		marks = append(marks, mark)
	}

	return marks, nil
}
