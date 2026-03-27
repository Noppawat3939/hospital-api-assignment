package route

import (
	"hospital-api/internal/client"
	"hospital-api/internal/handler"
	"hospital-api/internal/middleware"
	"hospital-api/internal/repository"
	"hospital-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPatientRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewPatientRepositroy(db)
	hosClient := client.NewHospitalClient("http://localhost:4000/mock")

	srv := service.NewPatientService(repo, *hosClient)

	h := handler.NewPatientHandler(srv)

	patient := r.Group("/patient")
	patient.Use(middleware.AuthGuard())
	{
		patient.POST("/search", h.PatientSearch)
	}
}
