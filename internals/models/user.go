package models

type User struct {
	Id             int    `json:"-"`
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	SecondName     string `json:"second_name"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	ProfileImgPath string `json:"profile_img_path"`
	Password       string `json:"password"`
}
