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

	group := r.Group("/staff")
	{
		group.POST("/create", h.StaffCreate)
		group.POST("/login", h.StaffLogin)
	}
}
