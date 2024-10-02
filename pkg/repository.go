package pkg

import (
	"context"

	"github.com/Filimonov-ua-d/to-do/models"
)

type Repository interface {
	GetUser(ctx context.Context, username, password string) (*models.User, error)
	Register(ctx context.Context, user models.User) error
	UserExist(ctx context.Context, username string) (bool, error)
	UpdateProfile(ctx context.Context, user *models.User) error
	ContactUs(ctx context.Context, contact models.Contact) error
}
