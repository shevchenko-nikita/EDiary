package models

type Assignment struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ClassId     int    `json:"class_id"`
	Statement   string `json:"statement"`
	TimeCreated string `json:"time_created"`
	DeadLine    string `json:"deadline"`
}
