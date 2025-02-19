package services

import (
	"database/sql"
	"fmt"
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

func DeleteClass(db *sql.DB, teacherId, classId int) error {
	actualTeacher, err := repository.GetClassById(db, classId)

	if err != nil {
		return err
	}

	if actualTeacher.TeacherId != teacherId {
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
