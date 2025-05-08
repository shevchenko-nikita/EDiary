package models

// this structure represents message in main page of class
type Message struct {
	ID         int    `json:"id"`
	ClassID    int    `json:"class_id"`
	UserID     int    `json:"user_id"`
	Text       string `json:"text"`
	TimePosted string `json:"time_posted"`
}

type ExpandedMessage struct {
	ID             int    `json:"id"`
	ClassID        int    `json:"class_id"`
	UserID         int    `json:"user_id"`
	UserName       string `json:"user_name"`
	UserProfileImg string `json:"user_profile_img"`
	Text           string `json:"text"`
	TimePosted     string `json:"time_posted"`
}
