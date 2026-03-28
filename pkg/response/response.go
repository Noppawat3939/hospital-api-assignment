package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrBodyInvalid       = "body invalid"
	ErrDataNotFound      = "data not found"
	ErrUnAuthorized      = "unauthorized"
	ErrMissingAuthHeader = "missing authorization header"
	ErrInvalidAuthFormat = "invalid authorization header format"
	ErrInvalidCrediental = "invalid creadiental"
	ErrHospitalInvalid   = "hospital invalid"
	ErrDupCreateStaff    = "staff already exits in this hospital"
)

func Success(c *gin.Context, data ...any) {
	res := gin.H{"success": true}

	if len(data) > 0 {
		res["data"] = data[0]
	}

	c.JSON(http.StatusOK, res)
}

func Error(c *gin.Context, status int, msg string, data ...any) {
	res := gin.H{"success": false}

	if msg != "" {
		res["message"] = msg
	}

	if len(data) > 0 {
		res["data"] = data
	}

	c.JSON(status, res)
}
