package repository

import (
	"hospital-api/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupHospitalTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory db: %v", err)
	}

	if err := db.AutoMigrate(&model.Hospital{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestHospitalRepository_FindOneByID(t *testing.T) {
	db := setupHospitalTestDB(t)
	repo := NewHospitalRepository(db)
	// insert a record
	hospital := model.Hospital{
		ID:   "HOSP1",
		Name: "Hospital One",
	}
	err := db.Create(&hospital).Error
	assert.NoError(t, err)

	// positive case
	found, err := repo.FindOneByID("HOSP1")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "Hospital One", found.Name)

	// negative case: not exist
	found, err = repo.FindOneByID("HOSP2")
	assert.Error(t, err)
	assert.Nil(t, found)
}
