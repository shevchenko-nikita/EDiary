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
	marksSum, err := repository.GetAllStudentGradesSum(db, userID)

	if err != nil {
		marksSum = 0.
	}

	if studentClasses > 0 {
		return marksSum / float32(studentClasses)
	}
	return 0.
}

func GetGradeDistribution(db *sql.DB, userID int) GradeDistribution {
	var distribution GradeDistribution
	distribution.Labels = []string{"незадовільно", "задовільно", "добре", "відмінно"}
	distribution.Data = []int{0, 0, 0, 0}

	grades, err := repository.GetAllStudentGrades(db, userID)
	if err != nil {
		grades = []int{}
	}

	for _, grade := range grades {
		if grade < 60 {
			distribution.Data[0]++
		} else if grade < 75 {
			distribution.Data[1]++
		} else if grade < 90 {
			distribution.Data[2]++
		} else {
			distribution.Data[3]++
		}
	}

	totalGradesAmount := len(grades)

	if totalGradesAmount != 0 {
		for i := range distribution.Data {
			distribution.Data[i] = distribution.Data[i] * 100 / totalGradesAmount
		}
	}

	return distribution
}

func GetStatisticInfo(db *sql.DB, userID int) (Statistic, error) {
	var statistic Statistic

	var studentInfo StudentInfo
	var subjects []models.ClassCard
	//var gradeDistribution GradeDistribution

	studentInfo.StudentClasses = GetStudentClassesNum(db, userID)
	studentInfo.TeachingClasses = GetTeachingClassesNum(db, userID)
	studentInfo.OverallAverage = GetOverallAverage(db, userID, studentInfo.StudentClasses)

	subjects, err := repository.GetEducationClasses(db, userID)
	if err != nil {
		subjects = []models.ClassCard{}
	}

	statistic.StudentInfo = studentInfo
	statistic.Subjects = subjects
	statistic.GradeDistribution = GetGradeDistribution(db, userID)

	return statistic, nil
}
