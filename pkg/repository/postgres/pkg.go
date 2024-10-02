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

func (p *PkgRepository) GetUser(ctx context.Context, email, password string) (u *models.User, err error) {

	user := new(User)

	selectSQL := "SELECT id, username, email, password_hash FROM users WHERE email=$1 AND password_hash=$2 LIMIT 1"
	err = p.DB.QueryRowContext(ctx, selectSQL, email, password).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return toModelUser(user), nil

}

func (p *PkgRepository) Register(ctx context.Context, user models.User) (err error) {
	_, err = p.DB.ExecContext(ctx, "INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3)", user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *PkgRepository) UserExist(ctx context.Context, username string) (bool, error) {
	var count int
	err := p.DB.GetContext(ctx, &count, "SELECT COUNT(*) FROM users WHERE username=$1", username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p *PkgRepository) UpdateProfile(ctx context.Context, user *models.User) error {
	_, err := p.DB.ExecContext(ctx, "UPDATE users SET username=$1, email=$2, image_url=$3 WHERE id=$4", user.Username, user.Email, user.ImageURL, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PkgRepository) ContactUs(ctx context.Context, contact models.Contact) error {
	_, err := p.DB.ExecContext(ctx, "INSERT INTO contact_us (name, email, message, phone, messanger, course) VALUES ($1, $2, $3, $4, $5, $6)",
		contact.Name, contact.Email, contact.Message, contact.Phone, contact.Messanger, contact.Course)
	if err != nil {
		return err
	}
	return nil
}
