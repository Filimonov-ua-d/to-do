package postgres

import "github.com/Filimonov-ua-d/to-do/models"

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password_hash"`
}

type Image struct {
	Id        int    `db:"id"`
	UserId    int    `db:"user_id"`
	ImagePath string `db:"image_path"`
	ImageUrl  string `db:"image_url"`
}

func toDBImage(i *models.Image) *Image {
	return &Image{
		Id:        i.Id,
		UserId:    i.UserId,
		ImagePath: i.ImagePath,
		ImageUrl:  i.ImageUrl,
	}
}

func toModelUser(u *User) *models.User {
	return &models.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}

func toModelImage(i *Image) *models.Image {
	return &models.Image{
		Id:        i.Id,
		UserId:    i.UserId,
		ImagePath: i.ImagePath,
		ImageUrl:  i.ImageUrl,
	}
}
