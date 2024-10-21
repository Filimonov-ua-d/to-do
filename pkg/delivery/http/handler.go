package http

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	useCase pkg.UseCase
}

func NewHandler(useCase pkg.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Login(c *gin.Context) {
	user := User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	res, token, err := h.useCase.Login(c.Request.Context(), user.Password, user.Email)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token, User: *res})
}

func (h *Handler) Register(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	token, err := h.useCase.Register(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, LoginResponse{Token: token, User: *user})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	if err := h.useCase.UpdateProfile(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{User: *user})
}

func (h *Handler) ContactUs(c *gin.Context) {
	form := models.Contact{}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	if err := h.useCase.ContactUs(c.Request.Context(), form); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) UploadVideo(c *gin.Context) {
	video := models.VideoLesson{}

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	if err := h.useCase.UploadVideo(c.Request.Context(), video); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

// TODO: Implement this method
func (h *Handler) UploadVideoFile(c *gin.Context) {
	// file, err := c.FormFile("file")
	// if err != nil {
	// c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
	// return
	// }

	// c.JSON(http.StatusOK, nil)
}

func (h *Handler) GetVideo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	video, err := h.useCase.GetVideo(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, video)
}

func (h *Handler) GetVideos(c *gin.Context) {
	videos, err := h.useCase.GetVideos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}

func (h *Handler) DeleteVideo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("course_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	if err := h.useCase.DeleteVideo(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) UploadPicture(c *gin.Context) {
	fileBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": err.Error()})
		return
	}

	if len(fileBytes) > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, ErrorResponse{"message": "File size exceeds the limit 2MB"})
		return
	}

	encodedFile, err := h.useCase.UploadPicture(c.Request.Context(), fileBytes, id)
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"encoded_file": encodedFile})
}
