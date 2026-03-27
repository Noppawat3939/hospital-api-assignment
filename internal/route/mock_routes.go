package route

import (
	"hospital-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMockRoutes(r *gin.Engine) {
	mock := r.Group("/mock")

	// prepare mock patients
	hospitalPatients := map[string][]gin.H{
		"3141c525-ca7d-4ae5-9d90-2a86ee5c8a2e": {
			{
				"first_name_th":  "สมหญิง",
				"middle_name_th": nil,
				"last_name_th":   "สวยงาม",
				"first_name_en":  "Somying",
				"middle_name_en": nil,
				"last_name_en":   "Suayngam",
				"date_of_birth":  "1995-05-10",
				"patient_hn":     "HN999999",
				"national_id":    "",
				"passport_id":    "passport123",
				"phone_number":   nil,
				"email":          nil,
				"gender":         "F",
			},
			{
				"first_name_th":  "สมชาย",
				"middle_name_th": nil,
				"last_name_th":   "ใจดี",
				"first_name_en":  "Somchai",
				"middle_name_en": nil,
				"last_name_en":   "Jaidee",
				"date_of_birth":  "1990-01-01",
				"patient_hn":     "HN123456",
				"national_id":    "national123",
				"passport_id":    "",
				"phone_number":   "0812345678",
				"email":          "somchai@example.com",
				"gender":         "M",
			},
		},
		"bc6f45eb-9b37-4939-bf46-11f7afc578b6": {
			{
				"first_name_th":  "บีบี",
				"middle_name_th": nil,
				"last_name_th":   "บีดี",
				"first_name_en":  "BB",
				"middle_name_en": nil,
				"last_name_en":   "BD",
				"date_of_birth":  "1988-08-08",
				"patient_hn":     "HN888888",
				"national_id":    "b123",
				"passport_id":    "",
				"phone_number":   "0811223344",
				"email":          "bb@example.com",
				"gender":         "F",
			},
		},
	}

	mock.GET("/hospital/:hospital/patient/search/:id", func(c *gin.Context) {
		hospital := c.Param("hospital")
		id := c.Param("id")

		patients, ok := hospitalPatients[hospital]
		if !ok {
			response.Error(c, http.StatusNotFound, response.ErrDataNotFound)
			return
		}

		for _, p := range patients {
			if p["national_id"] == id || p["passport_id"] == id {
				response.Success(c, p)
				return
			}
		}

		response.Error(c, http.StatusNotFound, response.ErrDataNotFound)
	})

}
