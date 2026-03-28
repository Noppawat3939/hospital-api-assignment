package route

import (
	"hospital-api/internal/mock"
	"hospital-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMockRoutes(r *gin.Engine) {
	group := r.Group("/mock")

	group.GET("/patient/search/:id", func(c *gin.Context) {
		id := c.Param("id")

		if id == "" {
			response.Error(c, http.StatusBadRequest, "id params invalid")
			return
		}

		patient, found := mock.FindPatientByID(id)

		if !found {
			response.Error(c, http.StatusNotFound, response.ErrDataNotFound)
			return
		}

		c.JSON(http.StatusOK, patient)
	})

}
