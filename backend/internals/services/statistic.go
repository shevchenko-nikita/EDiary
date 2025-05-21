package services

import (
	"database/sql"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
)

type Statistic struct {
	StudentInfo       StudentInfo       `json:"student_info"`
	Subjects          []Subject         `json:"subjects"`
	GradeDistribution GradeDistribution `json:"grade_distribution"`
}

type StudentInfo struct {
	OverallAverage  float32 `json:"overall_average"`
	StudentClasses  int     `json:"student_classes"`
	TeachingClasses int     `json:"teaching_classes"`
}

type Subject struct {
	SubjectID   int    `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Grade       int    `json:"grade"`
}

type GradeDistribution struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
	Colors []string `json:"colors"` // TBD
}

func GetStatisticInfo(db *sql.DB, userID int) (Statistic, error) {
	var statistic Statistic

	var studentInfo StudentInfo
	//var subjects []Subject

	studentClasses, err := repository.GetStudentClassesNum(db, userID)
	if err != nil {
		studentClasses = 0
	}
	studentInfo.StudentClasses = studentClasses

	teachingClasses, err := repository.GetTeachingClassesNum(db, userID)
	if err != nil {
		teachingClasses = 0
	}
	studentInfo.TeachingClasses = teachingClasses

	marksSum, err := repository.GetAllStudentMarks(db, userID)

	if err != nil {
		return Statistic{}, err
	}

	if studentClasses > 0 {
		studentInfo.OverallAverage = float32(marksSum) / float32(studentClasses)
	} else {
		studentInfo.OverallAverage = 0.
	}

	statistic.StudentInfo = studentInfo

	return statistic, nil
}
