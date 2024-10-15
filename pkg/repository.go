package pkg

import (
	"context"

	"github.com/Filimonov-ua-d/to-do/models"
)

type Repository interface {
	GetUser(ctx context.Context, email, password string) (*models.User, error)
	Register(ctx context.Context, user models.User) error
	UserExist(ctx context.Context, username string) (bool, error)
	UpdateProfile(ctx context.Context, user *models.User) error
	ContactUs(ctx context.Context, contact models.Contact) error
	UploadVideo(ctx context.Context, video models.VideoLesson) error
	GetVideo(ctx context.Context, id int) (*models.VideoLesson, error)
	GetVideos(ctx context.Context) ([]models.VideoLesson, error)
	DeleteVideo(ctx context.Context, id int) error
}
