package http

import (
	"errors"
	"net/http"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Handler struct {
	useCase pkg.UseCase
	logger  *zerolog.Logger
}

type LoginResponse struct {
	Token string `json:"token"`
}

type getImagesResponse struct {
	Images []*Image
}

func NewHandler(useCase pkg.UseCase, logger *zerolog.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *Handler) Login(c *gin.Context) {
	user := new(User)

	if err := c.ShouldBindJSON(&user); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "Login")

		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.useCase.Login(c.Request.Context(), user.Username, user.Password)
	if err != nil {
		if err == pkg.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h *Handler) UploadPicture(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	filename := file.Filename
	if err := c.SaveUploadedFile(file, "uploads/"+filename); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := c.MustGet(pkg.CtxUserKey).(*models.User)

	if err := h.useCase.UploadPicture(c.Request.Context(), user, filename); err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "UploadPicture")

		if errors.Is(err, pkg.ErrFileExist) {
			c.JSON(http.StatusOK, &UploadResponse{
				Error: ApiError{
					ErrorCode:        1,
					ErrorDescription: err.Error(),
				},
			})
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}

func (h *Handler) GetImages(c *gin.Context) {

	user := c.MustGet(pkg.CtxUserKey).(*models.User)

	im, err := h.useCase.GetImages(c.Request.Context(), user)
	if err != nil {

		h.logger.Error().
			Err(err).
			Str("Func:", "GetImages")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &getImagesResponse{
		Images: toModelImages(im),
	},
	)

}
