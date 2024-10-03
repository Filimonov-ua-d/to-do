package pkg

import (
	"context"

	"github.com/Filimonov-ua-d/to-do/models"
)

const CtxUserKey = "user"

type UseCase interface {
	Login(ctx context.Context, username, password, email string) (string, error)
	Register(ctx context.Context, user *models.User) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
	UpdateProfile(ctx context.Context, user *models.User) error
}
