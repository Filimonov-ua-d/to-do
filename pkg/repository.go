package pkg

import (
	"context"

	mod "github.com/Filimonov-ua-d/to-do/models"
)

type Repository interface {
	GetUser(ctx context.Context, username, password string) (*mod.User, error)
	ImageExists(ctx context.Context, filename string) (bool, error)
	UploadPicture(ctx context.Context, i *mod.Image) error
	GetImages(ctx context.Context) ([]*mod.Image, error)
}
