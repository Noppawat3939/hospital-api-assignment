package middleware

import (
	"hospital-api/pkg/jwt"
	"hospital-api/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const tokenPrefix = "Bearer "

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")

		statusCode := http.StatusUnauthorized

		if h == "" {
			response.Error(c, statusCode, response.ErrMissingAuthHeader)
			c.Abort()
			return
		}

		if !strings.HasPrefix(h, tokenPrefix) {
			response.Error(c, statusCode, response.ErrInvalidAuthFormat)
			c.Abort()
			return
		}

		// jwt validate
		jwtToken := strings.TrimPrefix(h, tokenPrefix)
		claims, err := jwt.Parse(jwtToken)

		if err != nil {
			response.Error(c, statusCode, response.ErrUnAuthorized)
			c.Abort()
			return
		}

		SetStaffContext(c, claims)

		c.Next()
	}
}
