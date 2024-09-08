package postgres

import (
	"context"
	"fmt"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/jmoiron/sqlx"
)

type PkgRepository struct {
	DB *sqlx.DB
}

func NewPkgRepository(db *sqlx.DB) *PkgRepository {
	return &PkgRepository{
		DB: db,
	}
}

func (pr *PkgRepository) GetUser(ctx context.Context, username, password string) (u *models.User, err error) {

	user := new(User)

	selectSQL := "SELECT id, username, password_hash FROM users WHERE username=$1 AND password_hash=$2 LIMIT 1"
	err = pr.DB.QueryRowContext(ctx, selectSQL, username, password).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return toModelUser(user), nil

}

func (pr *PkgRepository) CreateTask(ctx context.Context, t models.Task) (err error) {
	insertSQL := "INSERT INTO tasks (title, description, is_done) VALUES ($1, $2, $3)"
	_, err = pr.DB.ExecContext(ctx, insertSQL, t.Title, t.Description, t.IsDone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pr *PkgRepository) GetTasks(ctx context.Context) (tasks []*models.Task, err error) {
	rows, err := pr.DB.QueryxContext(ctx, "SELECT id, title, description, is_done FROM tasks")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		task := &models.Task{}
		err = rows.StructScan(task)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (pr *PkgRepository) GetTaskById(ctx context.Context, id int) (t *models.Task, err error) {
	task := new(models.Task)

	selectSQL := "SELECT id, title, description, is_done FROM tasks WHERE id=$1 LIMIT 1"
	err = pr.DB.QueryRowxContext(ctx, selectSQL, id).StructScan(task)
	if err != nil {
		fmt.Println(err)
		return
	}

	return task, nil
}

func (pr *PkgRepository) UpdateTask(ctx context.Context, t models.Task) (err error) {
	updateSQL := "UPDATE tasks SET title=$1, description=$2, is_done=$3 WHERE id=$4"
	_, err = pr.DB.ExecContext(ctx, updateSQL, t.Title, t.Description, t.IsDone, t.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (pr *PkgRepository) DeleteTask(ctx context.Context, id int) (err error) {
	deleteSQL := "DELETE FROM tasks WHERE id=$1"
	_, err = pr.DB.ExecContext(ctx, deleteSQL, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
