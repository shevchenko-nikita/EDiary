package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

func CreateNewAssignment(db *sql.DB, teacherId int, assignment *models.Assignment) error {
	class, err := repository.GetClassById(db, assignment.ClassID)

	if err != nil {
		return fmt.Errorf("class isn't exist")
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("user is not a teacher")
	}

	return repository.AddNewAssignment(db, assignment)
}

func DeleteAssignment(db *sql.DB, teacherId, assignmentId int) error {
	assignment, err := repository.GetAssignmentByID(db, assignmentId)

	if err != nil {
		return err
	}

	class, err := repository.GetClassById(db, assignment.ClassID)

	if err != nil {
		return err
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("user is not a teacher")
	}

	return repository.DeleteAssignment(db, assignmentId)
}

func GradeAssignment(db *sql.DB, teacherId int, mark models.Mark) error {
	assignment, err := repository.GetAssignmentByID(db, mark.AssignmentID)

	if err != nil {
		return err
	}

	class, err := repository.GetClassById(db, assignment.ClassID)

	if err != nil {
		return err
	}
	mark.ClassID = class.Id

	if class.TeacherId != teacherId {
		return fmt.Errorf("user is not a teacher")
	}

	if !repository.StudentExistInClass(db, mark.StudentID, class.Id) {
		return fmt.Errorf("user is not a student of the class")
	}

	if repository.MarkAlreadyExist(db, mark) {
		return repository.UpdateMark(db, mark)
	}

	return repository.AddNewMark(db, mark)
}

func GetAssignment(db *sql.DB, userID, assignmentID int) (models.Assignment, error) {
	assignment, err := repository.GetAssignmentByID(db, assignmentID)

	if err != nil {
		return models.Assignment{}, err
	}

	return assignment, nil
}

func GetAssignmentsList(db *sql.DB, userId, classId int) ([]models.Assignment, error) {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return nil, err
	}

	if class.TeacherId != userId && !repository.StudentExistInClass(db, userId, classId) {
		return nil, fmt.Errorf("user has no access")
	}

	return repository.GetAssignmentsList(db, classId)
}

func UpdateAssignment(db *sql.DB, teacherId int, newAssignmentInfo *models.Assignment) error {
	assignmentOrigin, err := repository.GetAssignmentByID(db, newAssignmentInfo.ID)

	if err != nil {
		return err
	}

	class, err := repository.GetClassById(db, assignmentOrigin.ClassID)

	if err != nil {
		return err
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("user doesn't have access")
	}

	return repository.UpdateAssignment(db, newAssignmentInfo)
}

func GetAllClassMarks(db *sql.DB, classId int) ([]models.Mark, error) {
	return repository.GetAllClassMarks(db, classId)
}

func GetMark(db *sql.DB, userID, assignmentID int) (float32, error) {
	return repository.GetMark(db, userID, assignmentID)
}
