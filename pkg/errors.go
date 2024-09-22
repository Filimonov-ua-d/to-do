package pkg

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrUserAlreadyExists  = errors.New("user already exists")
)
