package http

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ApiError struct {
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type ErrorResponse map[string]string

type LoginResponse struct {
	Token string `json:"token"`
}
