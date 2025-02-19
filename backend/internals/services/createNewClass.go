package services

import (
	"database/sql"
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
