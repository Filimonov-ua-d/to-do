package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username,omitempty"`
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
