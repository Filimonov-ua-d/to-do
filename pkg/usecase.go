package pkg

import (
	"context"

	"github.com/Filimonov-ua-d/to-do/models"
)

const CtxUserKey = "user"

type UseCase interface {
	Login(ctx context.Context, password, email string) (*models.User, string, error)
	Register(ctx context.Context, user *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
	UpdateProfile(ctx context.Context, user *models.User) error
	ContactUs(ctx context.Context, contact models.Contact) error
	UploadVideo(ctx context.Context, video models.VideoLesson) error
	GetVideo(ctx context.Context, id int) (*models.VideoLesson, error)
	GetVideos(ctx context.Context) ([]models.VideoLesson, error)
	DeleteVideo(ctx context.Context, id int) error
}
