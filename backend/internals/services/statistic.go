package services

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/models"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

type Statistic struct {
	StudentInfo       StudentInfo        `json:"student_info"`
	Subjects          []models.ClassCard `json:"subjects"`
	GradeDistribution GradeDistribution  `json:"grade_distribution"`
}

type StudentInfo struct {
	OverallAverage  float32 `json:"overall_average"`
	StudentClasses  int     `json:"student_classes"`
	TeachingClasses int     `json:"teaching_classes"`
}

type GradeDistribution struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
	Colors []string `json:"colors"` // TBD
}

func GetStudentClassesNum(db *sql.DB, userID int) int {
	studentClasses, err := repository.GetStudentClassesNum(db, userID)
	if err != nil {
		studentClasses = 0
	}

	return studentClasses
}

func GetTeachingClassesNum(db *sql.DB, userID int) int {
	teachingClasses, err := repository.GetTeachingClassesNum(db, userID)
	if err != nil {
		teachingClasses = 0
	}

	return teachingClasses
}

func GetOverallAverage(db *sql.DB, userID, studentClasses int) float32 {
	marksSum, err := repository.GetAllStudentMarks(db, userID)

	if err != nil {
		marksSum = 0
	}

	if studentClasses > 0 {
		return float32(marksSum) / float32(studentClasses)
	}
	return 0
}

func GetStatisticInfo(db *sql.DB, userID int) (Statistic, error) {
	var statistic Statistic

	var studentInfo StudentInfo
	var subjects []models.ClassCard
	var gradeDistribution GradeDistribution

	studentInfo.StudentClasses = GetStudentClassesNum(db, userID)
	studentInfo.TeachingClasses = GetTeachingClassesNum(db, userID)
	studentInfo.OverallAverage = GetOverallAverage(db, userID, studentInfo.StudentClasses)

	subjects, err = repository.GetEducationClasses(db, userID)
	if err != nil {
		subjects = []models.ClassCard{}
	}

	statistic.StudentInfo = studentInfo
	statistic.Subjects = subjects

	return statistic, nil
}
