package models

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

type Contact struct {
	Name      string `json:"name"`
	Phone     string `json:"phone,omitempty"`
	Messanger string `json:"messanger,omitempty"`
	Email     string `json:"email"`
	Message   string `json:"message,omitempty"`
	Course    string `json:"course"`
}

type VideoLesson struct {
	Id       int    `json:"id,omitempty"`
	CourseID int    `json:"course_id"`
	URL      string `json:"url"`
	Comment  string `json:"comment,omitempty"`
}
