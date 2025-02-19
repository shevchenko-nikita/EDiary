package models

type Class struct {
	CodeId    string `json:"class_code_id"`
	Name      string `json:"name"`
	TeacherId int    `json:"teacher_id"`
}
