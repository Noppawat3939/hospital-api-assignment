package migration

import (
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

func MigrationCreateHospital(db *gorm.DB) error {
	return db.Migrator().AutoMigrate(&model.Hospital{})
}
