package pkg

import (
	"context"

	mod "github.com/Filimonov-ua-d/to-do/models"
)

const CtxUserKey = "user"

type UseCase interface {
	Login(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*mod.User, error)
	CreateTask(ctx context.Context, t mod.Task) (err error)
	GetTasks(ctx context.Context) ([]*mod.Task, error)
	GetTaskById(ctx context.Context, id int) (*mod.Task, error)
	UpdateTask(ctx context.Context, t mod.Task) (err error)
	DeleteTask(ctx context.Context, id int) (err error)
}
