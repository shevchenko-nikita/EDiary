package models

type UserFile struct {
	ID           int    `json:id`
	AssignmentId int    `json:assignment_id`
	StudentId    int    `json:student_id`
	Name         string `json:file_name`
	Path         string `json:path`
}
