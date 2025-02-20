package models

type Mark struct {
	Id           int     `json:"id"`
	AssignmentId int     `json:"assignment_id"`
	StudentId    int     `json:"student_id"`
	Mark         float32 `json:"mark"`
}
