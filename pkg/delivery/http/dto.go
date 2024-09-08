package http

import (
	"github.com/Filimonov-ua-d/to-do/models"
)

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Image struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	ImagePath string `json:"image_path"`
	ImageUrl  string `json:"image_url"`
}

type ApiError struct {
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

type UploadResponse struct {
	Error ApiError `json:"error"`
}

func toModelUser(u *User) *models.User {
	return &models.User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
}

func toImage(i *models.Image) *Image {
	return &Image{
		Id:        i.Id,
		UserId:    i.UserId,
		ImagePath: i.ImagePath,
		ImageUrl:  i.ImageUrl,
	}
}

func toModelImages(ims []*models.Image) []*Image {

	var images []*Image

	for _, v := range ims {
		images = append(images, (toImage)(v))
	}

	return images
}
