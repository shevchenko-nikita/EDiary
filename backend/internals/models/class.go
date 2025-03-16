package models

type Class struct {
	Id        int    `json:"id"`
	Code      string `json:"class_code"`
	Name      string `json:"name"`
	TeacherId int    `json:"teacher_id"`
}

type ClassCard struct {
	Id             int     `json:"class_id"`
	Code           string  `json:"class_code"`
	Name           string  `json:"name"`
	FirstName      string  `json:"first_name"`
	MiddleName     string  `json:"middle_name"`
	SecondName     string  `json:"second_name"`
	ProfileImgPath string  `json:"profile_img_path"`
	Grade          float32 `json:"grade"`
}
