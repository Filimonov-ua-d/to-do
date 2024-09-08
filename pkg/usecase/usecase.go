package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/rs/zerolog"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type PkgUseCase struct {
	PkgRepo        pkg.Repository
	logger         *zerolog.Logger
	signingKey     []byte
	expireDuration time.Duration
	hashSalt       string
	port           string
}

func NewPkgUseCase(PkgRepo pkg.Repository,
	logger *zerolog.Logger,
	signingKey []byte,
	hashSalt string,
	tokenTTLSeconds time.Duration,
	port string) *PkgUseCase {
	return &PkgUseCase{
		PkgRepo:        PkgRepo,
		logger:         logger,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
		hashSalt:       hashSalt,
		port:           port,
	}
}

func (p *PkgUseCase) Login(ctx context.Context, username, password string) (string, error) {

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(p.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := p.PkgRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", pkg.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(p.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(p.signingKey)
}

func (p *PkgUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {

	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return p.signingKey, nil
	})

	if err != nil {

		p.logger.Error().
			Err(err).
			Str("Func:", "ParseToken")

		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, pkg.ErrInvalidAccessToken
}

func (p *PkgUseCase) UploadPicture(ctx context.Context, u *models.User, filename string) error {
	dbUser, err := p.PkgRepo.GetUser(ctx, u.Username, u.Password)
	if err != nil {
		p.logger.Error().
			Err(err).
			Str("Func:", "UploadPicture")

		return err
	}

	exists, err := p.PkgRepo.ImageExists(ctx, "uploads/"+filename)
	if err != nil {
		p.logger.Error().
			Err(err).
			Str("Func:", "UploadPicture")

		return err
	}

	if exists {
		return fmt.Errorf("Error insert file %s.%w", filename, pkg.ErrFileExist)
	}

	im := &models.Image{
		UserId:    dbUser.Id,
		ImagePath: "uploads/" + filename,
		ImageUrl:  "http://localhost:" + p.port + "/uploads/" + filename,
	}

	err = p.PkgRepo.UploadPicture(ctx, im)

	p.logger.Error().
		Err(err).
		Str("Func:", "UploadPicture")

	return err
}

func (p *PkgUseCase) GetImages(ctx context.Context, u *models.User) ([]*models.Image, error) {
	c, err := p.PkgRepo.GetImages(ctx)

	p.logger.Error().
		Err(err).
		Str("Func:", "GetImages")

	return c, err
}
