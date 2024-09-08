package models

type User struct {
	Id       int
	Username string
	Password string
}

type Image struct {
	Id        int
	UserId    int
	ImagePath string
	ImageUrl  string
}
