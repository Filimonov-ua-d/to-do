package postgres

import (
	"context"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type PkgRepository struct {
	DB     *sqlx.DB
	logger *zerolog.Logger
}

func NewPkgRepository(db *sqlx.DB, logger *zerolog.Logger) *PkgRepository {
	return &PkgRepository{
		DB:     db,
		logger: logger,
	}
}

func (pr *PkgRepository) GetUser(ctx context.Context, username, password string) (u *models.User, err error) {

	user := new(User)

	selectSQL := "SELECT id, username, password_hash FROM users WHERE username=$1 AND password_hash=$2 LIMIT 1"
	err = pr.DB.QueryRowContext(ctx, selectSQL, username, password).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {

		pr.logger.Error().
			Err(err).
			Str("Func:", "GetUser")

		return nil, err
	}

	return toModelUser(user), nil

}

func (pr *PkgRepository) ImageExists(c context.Context, filename string) (bool, error) {

	var count int
	err := pr.DB.QueryRowxContext(c, "SELECT COUNT (*) FROM images WHERE image_path = $1", filename).Scan(&count)

	if err != nil {
		pr.logger.Error().
			Err(err).
			Str("Func:", "ImageExists")
		return false, err
	}

	return count > 0, nil
}

func (pr *PkgRepository) UploadPicture(c context.Context, i *models.Image) (err error) {

	insertSQL := "insert into images (user_id,image_path,image_url) VALUES (:user_id,:image_path,:image_url)"

	var image = toDBImage(i)

	_, err = pr.DB.NamedExec(insertSQL, image)
	if err != nil {

		pr.logger.Error().
			Err(err).
			Str("Func:", "UploadPicture")

		return
	}

	return

}

func (pr *PkgRepository) GetImages(ctx context.Context) (im []*models.Image, err error) {

	selectSQL := "SELECT * FROM images"

	var images []*Image

	if err = pr.DB.Select(&images, selectSQL); err != nil {

		pr.logger.Error().
			Err(err).
			Str("Func:", "GetImages")

		return
	}

	for _, v := range images {
		im = append(im, toModelImage(v))

	}

	return

}
