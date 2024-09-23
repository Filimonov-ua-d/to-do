package postgres

import "github.com/Filimonov-ua-d/to-do/models"

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password_hash"`
}

func toModelUser(u *User) *models.User {
	return &models.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}
