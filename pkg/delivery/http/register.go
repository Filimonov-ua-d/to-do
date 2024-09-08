package http

import (
	"net/http"
	"os"

	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RegisterHTTPEndpoints(router *gin.Engine, pk pkg.UseCase) {

	loggerHandler := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("Layer:", "handler").
		Str("Service:", "test_task").
		Logger()

	h := NewHandler(pk, &loggerHandler)

	router.StaticFS("/uploads", http.Dir("uploads"))

	router.POST("/login", h.Login)

	mw := NewAuthMiddleware(pk)
	apiRouter := router.Group("/api", mw)

	apiRouter.POST("/upload-picture", h.UploadPicture)
	apiRouter.GET("/images", h.GetImages)
}
