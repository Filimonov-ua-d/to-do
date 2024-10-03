package http

import (
	"net/http"
	"strings"

	"github.com/Filimonov-ua-d/to-do/pkg"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	usecase pkg.UseCase
}

func NewAuthMiddleware(usecase pkg.UseCase) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle

}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == pkg.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}
		c.AbortWithStatus(status)
		return
	}

	c.Set(pkg.CtxUserKey, user)
}
