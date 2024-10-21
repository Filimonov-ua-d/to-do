package http

import (
	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, pk pkg.UseCase) {
	h := NewHandler(pk)

	router.POST("/login", h.Login)
	router.POST("/register", h.Register)
	router.POST("/contact-us", h.ContactUs)

	mw := NewAuthMiddleware(pk)
	apiRouter := router.Group("/api", mw)

	apiRouter.PUT("/update-profile", h.UpdateProfile)
	apiRouter.POST("/upload-video", h.UploadVideo)
	apiRouter.GET("/videos", h.GetVideos)
	apiRouter.GET("/video/:course_id", h.GetVideo)
	apiRouter.DELETE("/video/:course_id", h.DeleteVideo)
	apiRouter.PUT("/upload-image/:id", h.UploadPicture)
}
