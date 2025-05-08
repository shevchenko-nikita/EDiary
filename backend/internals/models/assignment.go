package models

type Assignment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ClassID     int    `json:"class_id"`
	Statement   string `json:"statement"`
	TimeCreated string `json:"time_created"`
	DeadLine    string `json:"deadline"`
}
