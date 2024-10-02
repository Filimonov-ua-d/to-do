package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}
