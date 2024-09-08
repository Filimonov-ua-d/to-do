package http

import (
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, pk pkg.UseCase) {
	h := NewHandler(pk)

	router.POST("/login", h.Login)

	mw := NewAuthMiddleware(pk)
	apiRouter := router.Group("/api", mw)

	// apiRouter.POST("/upload-picture", h.UploadPicture)
	// apiRouter.GET("/images", h.GetImages)
	apiRouter.POST("/create-task", h.CreateTask)
	apiRouter.GET("/tasks", h.GetTasks)
	apiRouter.GET("/tasks/:id", h.GetTaskById)
	apiRouter.PUT("/tasks/:id", h.UpdateTask)
	apiRouter.DELETE("/tasks/:id", h.DeleteTask)
}
