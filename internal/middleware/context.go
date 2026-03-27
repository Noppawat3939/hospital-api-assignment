package middleware

import (
	"hospital-api/internal/model"
	"hospital-api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const staffContext = "staff"

func SetStaffContext(c *gin.Context, claims *jwt.Claims) {
	staff := &model.Staff{
		HospitalID: claims.Hospital,
		Username:   claims.Username,
	}

	c.Set(staffContext, staff)
}

func GetStaffContext(c *gin.Context) (*model.Staff, bool) {
	val, exists := c.Get(staffContext)

	if !exists {
		return nil, false
	}

	staff, ok := val.(*model.Staff)
	if !ok {
		return nil, false
	}

	return staff, true
}
