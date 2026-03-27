package route

import (
	"hospital-api/internal/handler"
	"hospital-api/internal/repository"
	"hospital-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterStaffRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewStaffRepository(db)
	hosRepo := repository.NewHospitalRepository(db)

	srv := service.NewStaffService(repo)
	hosSrv := service.NewHospitalService(hosRepo)

	h := handler.NewStaffHandler(srv, hosSrv)

	staff := r.Group("/staff")
	{
		staff.POST("/create", h.StaffCreate)
		staff.POST("/login", h.StaffLogin)
	}
}
