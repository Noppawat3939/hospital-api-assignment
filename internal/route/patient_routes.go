package route

import (
	"hospital-api/internal/client"
	"hospital-api/internal/handler"
	"hospital-api/internal/middleware"
	"hospital-api/internal/repository"
	"hospital-api/internal/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPatientRoutes(r *gin.Engine, db *gorm.DB) {
	hisBaseURL := os.Getenv("HIS_BASE_URL")
	if hisBaseURL == "" {
		panic("HIS_BASE_URL is required")
	}

	repo := repository.NewPatientRepositroy(db)
	hosClient := client.NewHospitalClient(hisBaseURL)

	srv := service.NewPatientService(repo, *hosClient)

	h := handler.NewPatientHandler(srv)

	group := r.Group("/patient")
	group.Use(middleware.AuthGuard())
	{
		group.POST("/search", h.PatientSearch)
	}
}
