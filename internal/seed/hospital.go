package seed

import (
	"hospital-api/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedHospital(db *gorm.DB) error {
	var count int64
	hospitals := []model.Hospital{
		{ID: uuid.NewString(), Name: "Hospital A"},
		{ID: uuid.NewString(), Name: "Hospital B"},
		{ID: uuid.NewString(), Name: "Hospital C"},
	}

	if err := db.Model(&model.Hospital{}).Count(&count).Error; err != nil {
		return err
	}

	if int(count) == len(hospitals) {
		return nil
	}

	for _, h := range hospitals {
		if err := db.Create(&h).Error; err != nil {
			return err
		}
	}
	return nil
}
