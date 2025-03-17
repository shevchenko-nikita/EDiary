package models

type Mark struct {
	Id           int     `json:"id"`
	ClassId      int     `json:"class_id"`
	AssignmentId int     `json:"assignment_id"`
	StudentId    int     `json:"student_id"`
	Mark         float32 `json:"mark"`
}
