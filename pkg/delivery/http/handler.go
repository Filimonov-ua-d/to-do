package http

import (
	"fmt"
	"net/http"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
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

	token, err := h.useCase.Login(c.Request.Context(), user.Username, user.Password, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token})
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

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
