package repository

import (
	"hospital-api/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupStaffTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory db: %v", err)
	}

	if err := db.AutoMigrate(&model.Staff{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestStaffRepository_Create(t *testing.T) {
	db := setupStaffTestDB(t)
	repo := NewStaffRepository(db)

	staff := model.Staff{
		Username:   "john_doe",
		Password:   "password123",
		HospitalID: "HOSP1",
	}

	created, err := repo.Create(staff)
	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, "john_doe", created.Username)
	assert.Equal(t, "HOSP1", created.HospitalID)
}

func TestStaffRepository_FindOneByUsernameAndHospitalID(t *testing.T) {
	db := setupStaffTestDB(t)
	repo := NewStaffRepository(db)

	// insert a record
	staff := model.Staff{
		Username:   "alice",
		Password:   "pass",
		HospitalID: "HOSP1",
	}
	_, err := repo.Create(staff)
	assert.NoError(t, err)

	// positive case
	found, err := repo.FindOneByUsernameAndHospitalID("alice", "HOSP1")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "alice", found.Username)
	assert.Equal(t, "HOSP1", found.HospitalID)

	// negative case: wrong hospital
	found, err = repo.FindOneByUsernameAndHospitalID("alice", "HOSP2")
	assert.Error(t, err)
	assert.Nil(t, found)

	// negative case: wrong username
	found, err = repo.FindOneByUsernameAndHospitalID("bob", "HOSP1")
	assert.Error(t, err)
	assert.Nil(t, found)
}
