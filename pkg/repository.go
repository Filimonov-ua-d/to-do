package pkg

import (
	"context"

	"github.com/Filimonov-ua-d/to-do/models"
)

type Repository interface {
	GetUser(ctx context.Context, username, password string) (*models.User, error)
	CreateTask(ctx context.Context, t models.Task) (err error)
	GetTasks(ctx context.Context) ([]*models.Task, error)
	GetTaskById(ctx context.Context, id int) (*models.Task, error)
	UpdateTask(ctx context.Context, t models.Task) (err error)
	DeleteTask(ctx context.Context, id int) (err error)
}
