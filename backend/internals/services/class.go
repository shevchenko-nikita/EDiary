package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"math/rand"
)

const CHARSET = "abcdfghjklmnpqrstvwxyz" +
	"ABCDFGHJKLMNPQRSTVWXYZ" +
	"0123456789" +
	"#$@"

const CODE_LEN int8 = 7

func generateClassCode() string {
	b := make([]byte, CODE_LEN)

	for i := 0; i < 7; i++ {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}

	return string(b)
}

func CreateNewClass(db *sql.DB, teacherId int, className string) error {
	var classCode string
	for true {
		classCode = generateClassCode()
		if !repository.ClassExists(db, classCode) {
			break
		}
	}

	return repository.CreateNewClass(db, classCode, className, teacherId)
}

func JoinTheClass(db *sql.DB, studentId int, classCode string) error {
	if !repository.ClassExists(db, classCode) {
		return fmt.Errorf("class %s does not exist", classCode)
	}

	class, err := repository.GetClassByCode(db, classCode)

	if err != nil {
		return err
	}

	if class.TeacherId == studentId {
		return fmt.Errorf("User is a teacher of the class")
	}

	return repository.JoinTheClass(db, studentId, class.Id)
}

func UpdateClass(db *sql.DB, teacherId, classId int, newClassName string) error {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return err
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("User is not a teacher of the class")
	}

	return repository.UpdateClass(db, classId, newClassName)
}

func DeleteClass(db *sql.DB, teacherId, classId int) error {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return err
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("you are not owner of the class")
	}

	return repository.DeleteClass(db, classId)
}

func LeaveClass(db *sql.DB, studentId, classId int) error {
	//if !repository.StudentExistInClass(db, studentId, classId) {
	//	return fmt.Errorf("the student does not exist in the class")
	//}

	return repository.LeaveClass(db, studentId, classId)
}

func GetUsersList(db *sql.DB, userId, classId int) ([]models.User, error) {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return nil, err
	}

	if !repository.StudentExistInClass(db, userId, classId) &&
		class.TeacherId != userId {
		return nil, fmt.Errorf("User has not access")
	}

	return repository.GetUsersList(db, classId)
}

func CreateNewAssignment(db *sql.DB, teacherId int, assignment *models.Assignment) error {
	class, err := repository.GetClassById(db, assignment.ClassId)

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

	class, err := repository.GetClassById(db, assignment.ClassId)

	if err != nil {
		return err
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("user is not a teacher")
	}

	return repository.DeleteAssignment(db, assignmentId)
}

func GradeAssignment(db *sql.DB, teacherId int, mark models.Mark) error {
	assignment, err := repository.GetAssignmentByID(db, mark.AssignmentId)

	if err != nil {
		return err
	}

	class, err := repository.GetClassById(db, assignment.ClassId)

	if err != nil {
		return err
	}

	if class.TeacherId != teacherId {
		return fmt.Errorf("user is not a teacher")
	}

	if !repository.StudentExistInClass(db, mark.StudentId, class.Id) {
		return fmt.Errorf("user is not a student of the class")
	}

	if repository.MarkAlreadyExist(db, mark) {
		return repository.UpdateMark(db, mark)
	}

	return repository.AddNewMark(db, mark)
}
