package repository

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/pkg/common"
	"hospital-api/pkg/timeutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupPatientTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}

	if err := db.AutoMigrate(&model.Patient{}); err != nil {
		t.Fatalf("failed to migrate Patient: %v", err)
	}

	return db
}

func TestPatientRepository_FindAll(t *testing.T) {
	db := setupPatientTestDB(t)
	repo := NewPatientRepositroy(db)

	// insert mock data
	patients := []model.Patient{
		{PatientHN: "HN1", FirstNameTH: "สมชาย", FirstNameEN: "Somchai", HospitalID: "HOSP1", NationalID: common.StringToPtr("111"), PassportID: common.StringToPtr("P111")},
		{PatientHN: "HN2", FirstNameTH: "สมหญิง", FirstNameEN: "Somying", HospitalID: "HOSP1", NationalID: common.StringToPtr("222"), PassportID: common.StringToPtr("P222")},
		{PatientHN: "HN3", FirstNameTH: "สมหมาย", FirstNameEN: "Sommai", HospitalID: "HOSP2", NationalID: common.StringToPtr("333"), PassportID: common.StringToPtr("P333")},
	}

	for _, p := range patients {
		_, err := repo.Create(p)
		assert.NoError(t, err)
	}

	// positive case: filter by hospital
	req := dto.SearchPatientRequest{Limit: nil, Page: nil}
	results, err := repo.FindAll("HOSP1", req)
	assert.NoError(t, err)
	assert.Len(t, results, 2)

	// filter by national_id
	nid := "111"
	req = dto.SearchPatientRequest{NationalID: &nid}
	results, err = repo.FindAll("HOSP1", req)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "HN1", results[0].PatientHN)

	// negative case: no match
	nid = "999"
	req = dto.SearchPatientRequest{NationalID: &nid}
	results, err = repo.FindAll("HOSP1", req)
	assert.NoError(t, err)
	assert.Len(t, results, 0)
}

func TestPatientRepository_FindOneByIdentity(t *testing.T) {
	db := setupPatientTestDB(t)
	repo := NewPatientRepositroy(db)

	p := model.Patient{
		PatientHN:  "HN100",
		NationalID: common.StringToPtr("1234567890"),
		PassportID: common.StringToPtr("A9876543"),
		HospitalID: "HOSP1",
	}
	_, err := repo.Create(p)
	assert.NoError(t, err)

	// positive case: by national_id
	nid := "1234567890"
	req := dto.SearchPatientRequest{NationalID: &nid}
	found, err := repo.FindOneByIdentity(req)
	assert.NoError(t, err)
	assert.Equal(t, "HN100", found.PatientHN)

	// positive case: by passport_id
	pid := "A9876543"
	req = dto.SearchPatientRequest{PassportID: &pid}
	found, err = repo.FindOneByIdentity(req)
	assert.NoError(t, err)
	assert.Equal(t, "HN100", found.PatientHN)

	// negative case: not found
	nid = "0000000"
	req = dto.SearchPatientRequest{NationalID: &nid}
	found, err = repo.FindOneByIdentity(req)
	assert.Error(t, err)
	assert.Nil(t, found.PatientHN)
}

func TestPatientRepository_Create(t *testing.T) {
	db := setupPatientTestDB(t)
	repo := NewPatientRepositroy(db)

	dob, err := timeutil.ParseDate("1990-01-01")
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	patient := model.Patient{
		FirstNameTH: "สมชาย",
		LastNameTH:  "ใจดี",
		FirstNameEN: "Somchai",
		LastNameEN:  "Jaidee",
		DateOfBirth: dob,
		PatientHN:   "HN001",
		NationalID:  common.StringToPtr("1234567890123"),
		PassportID:  common.StringToPtr("A1234567"),
		HospitalID:  "HOSP1",
	}

	created, err := repo.Create(patient)
	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, "HN001", created.PatientHN)
}
