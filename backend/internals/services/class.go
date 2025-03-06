package services

import (
	"database/sql"
	"fmt"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

const CODE_LEN int8 = 7

func generateClassCode() string {
	return generateCode(CODE_LEN)
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

func GetStudentsList(db *sql.DB, userId, classId int) ([]models.User, error) {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return nil, err
	}

	if !repository.StudentExistInClass(db, userId, classId) &&
		class.TeacherId != userId {
		return nil, fmt.Errorf("User has not access")
	}

	return repository.GetStudentsList(db, classId)
}

func GetClassTeacher(db *sql.DB, userId, classId int) (models.User, error) {
	class, err := repository.GetClassById(db, classId)

	if err != nil {
		return models.User{}, err
	}

	if !repository.StudentExistInClass(db, userId, classId) &&
		class.TeacherId != userId {
		return models.User{}, fmt.Errorf("User has not access")
	}

	return repository.GetUserById(db, class.TeacherId)
}
