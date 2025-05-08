package models

type Mark struct {
	ID           int     `json:"id"`
	ClassID      int     `json:"class_id"`
	AssignmentID int     `json:"assignment_id"`
	StudentID    int     `json:"student_id"`
	Mark         float32 `json:"mark"`
}
