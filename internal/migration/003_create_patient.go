package migration

import (
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

func MigrationCreatePatientGenderEnum(db *gorm.DB) error {
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'patient_gender')").Scan(&exists).Error
	if err != nil {
		return err
	}
	if !exists {
		return db.Exec("CREATE TYPE patient_gender AS ENUM ('M', 'F')").Error
	}
	return nil
}

func MigrationCreatePatient(db *gorm.DB) error {
	// create enum first
	if err := MigrationCreatePatientGenderEnum(db); err != nil {
		return err
	}

	if err := db.Migrator().AutoMigrate(&model.Patient{}); err != nil {
		return err
	}

	if !db.Migrator().HasIndex(&model.Patient{}, "idx_patient_hospital_th") {
		db.Migrator().CreateIndex(&model.Patient{}, "hospital_id_firstname_th_middle_name_th_last_name_th")
	}
	if !db.Migrator().HasIndex(&model.Patient{}, "idx_patient_hospital_en") {
		db.Migrator().CreateIndex(&model.Patient{}, "hospital_id_firstname_en_middle_name_en_last_name_en")
	}

	return nil
}
