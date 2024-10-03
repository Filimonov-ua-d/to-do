package http

import (
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, pk pkg.UseCase) {
	h := NewHandler(pk)

	router.POST("/login", h.Login)
	router.POST("/register", h.Register)

	mw := NewAuthMiddleware(pk)
	apiRouter := router.Group("/api", mw)

	apiRouter.PUT("/update-profile", h.UpdateProfile)
	// apiRouter.POST("/create-task", h.CreateTask)
}
