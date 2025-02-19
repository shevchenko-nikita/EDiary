package models

type Class struct {
	Id        int    `json:"id"`
	Code      string `json:"class_code"`
	Name      string `json:"name"`
	TeacherId int    `json:"teacher_id"`
}
