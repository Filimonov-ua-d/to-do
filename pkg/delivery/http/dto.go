package http

import "github.com/Filimonov-ua-d/to-do/models"

type User struct {
	Id       int    `json:"-,omitempty"`
	Username string `json:"name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ImageURL string `json:"image_url,omitempty"`
}

type ApiError struct {
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type ErrorResponse map[string]string

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
