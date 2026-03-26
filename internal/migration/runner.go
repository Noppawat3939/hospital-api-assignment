package migration

import (
	"log"

	"gorm.io/gorm"
)

var migrations = map[string]func(*gorm.DB) error{
	"001_create_hospital": MigrationCreateHospital,
	"002_create_staff":    MigrationCreateStaff,
	"003_create_patient":  MigrationCreatePatient,
}

func RunMigration(db *gorm.DB, fileName string) {
	if fn, ok := migrations[fileName]; ok {
		if err := fn(db); err != nil {
			log.Fatalf("migration %s failed: %v", fileName, err)
		}
		log.Printf("migration %s completed", fileName)
	} else {
		log.Fatalf("unknow migration file %s", fileName)
	}
}
