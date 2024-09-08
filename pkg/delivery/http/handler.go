package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Filimonov-ua-d/to-do/models"
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase pkg.UseCase
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewHandler(useCase pkg.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Login(c *gin.Context) {
	user := new(User)

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
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

func (h *Handler) CreateTask(c *gin.Context) {
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.String(http.StatusBadRequest, err.Error(), err)
		return
	}

	err := h.useCase.CreateTask(c.Request.Context(), task)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetTasks(c *gin.Context) {
	tasks, err := h.useCase.GetTasks(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	task, err := h.useCase.GetTaskById(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err := h.useCase.UpdateTask(c.Request.Context(), task)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.useCase.DeleteTask(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), err)
		return
	}

	c.Status(http.StatusOK)
}
