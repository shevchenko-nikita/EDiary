package models

// this structure represents message in main page of class
type Message struct {
	Id         int    `json:"id"`
	ClassId    int    `json:"class_id"`
	UserId     int    `json:"user_id"`
	Text       string `json:"text"`
	TimePosted string `json:"time_posted"`
}
