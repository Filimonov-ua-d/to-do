package usecase

import (
	"bytes"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	accessKey      string
	secretKey      string
	region         string
	bucket         string
}

func NewPkgUseCase(PkgRepo pkg.Repository,
	signingKey []byte,
	hashSalt string,
	tokenTTL time.Duration,
	port string,
	ak string, sk string, region string, bucket string,
) *PkgUseCase {
	return &PkgUseCase{
		PkgRepo:        PkgRepo,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
		hashSalt:       hashSalt,
		port:           port,
		accessKey:      ak,
		secretKey:      sk,
		region:         region,
		bucket:         bucket,
	}
}

func (p *PkgUseCase) Login(ctx context.Context, password, email string) (*models.User, string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(p.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := p.PkgRepo.GetUser(ctx, email, password)
	if err != nil {
		return nil, "", err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(p.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	res, err := token.SignedString(p.signingKey)
	return user, res, err
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

func (p *PkgUseCase) UploadPicture(ctx context.Context, fileBytes []byte, filename, fileExtension string, fileSize, userID int64) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(p.region),
		Credentials: credentials.NewStaticCredentials(p.accessKey, p.secretKey, ""),
	})
	if err != nil {
		return "", fmt.Errorf("error creating session: %w", err)
	}

	svc := s3.New(sess)

	fileType := fmt.Sprintf("image/%s", strings.TrimPrefix(fileExtension, "."))
	key := fmt.Sprintf("uploads/%d/%s", userID, filename)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(p.bucket),
		Key:                aws.String(key),
		Body:               bytes.NewReader(fileBytes),
		ContentLength:      aws.Int64(fileSize),
		ContentType:        aws.String(fileType),
		ContentDisposition: aws.String("inline"),
	})
	if err != nil {
		log.Fatalf("Не вдалося завантажити файл: %v", err)
	}

	resUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", p.bucket, p.region, key)

	exists, err := p.PkgRepo.ImageExists(ctx, resUrl, int(userID))
	if err != nil {
		return "", err
	}

	if exists {
		return "", errors.New("error insert file. File already exists")
	}

	if err = p.PkgRepo.UploadPicture(ctx, resUrl, int(userID)); err != nil {
		return "", err
	}

	return resUrl, err
}
