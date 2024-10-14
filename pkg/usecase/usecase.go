package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type PkgUseCase struct {
	PkgRepo        pkg.Repository
	signingKey     []byte
	expireDuration time.Duration
	hashSalt       string
	port           string
}

func NewPkgUseCase(PkgRepo pkg.Repository,
	signingKey []byte,
	hashSalt string,
	tokenTTLSeconds time.Duration,
	port string) *PkgUseCase {
	return &PkgUseCase{
		PkgRepo:        PkgRepo,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
		hashSalt:       hashSalt,
		port:           port,
	}
}

func (p *PkgUseCase) Login(ctx context.Context, username, password, email string) (string, error) {
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
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return p.signingKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, pkg.ErrInvalidAccessToken
}

func (p *PkgUseCase) Register(ctx context.Context, user *models.User) (string, error) {

	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(p.hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	exists, err := p.PkgRepo.UserExist(ctx, user.Username)
	if err != nil {
		return "", err
	}

	if exists {
		return "", pkg.ErrUserAlreadyExists
	}

	err = p.PkgRepo.Register(ctx, *user)
	if err != nil {
		return "", err
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

func (p *PkgUseCase) UpdateProfile(ctx context.Context, user *models.User) error {
	return p.PkgRepo.UpdateProfile(ctx, user)
}

func (p *PkgUseCase) ContactUs(ctx context.Context, contact models.Contact) error {
	return p.PkgRepo.ContactUs(ctx, contact)
}

func (p *PkgUseCase) UploadVideo(ctx context.Context, video models.VideoLesson) error {
	return p.PkgRepo.UploadVideo(ctx, video)
}

func (p *PkgUseCase) GetVideos(ctx context.Context) ([]models.VideoLesson, error) {
	return p.PkgRepo.GetVideos(ctx)
}

func (p *PkgUseCase) GetVideo(ctx context.Context, id int) (*models.VideoLesson, error) {
	return p.PkgRepo.GetVideo(ctx, id)
}

func (p *PkgUseCase) DeleteVideo(ctx context.Context, id int) error {
	return p.PkgRepo.DeleteVideo(ctx, id)
}
