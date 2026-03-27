package repository

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

type PatientRepository interface {
	FindAll(hospitalID string, req dto.SearchPatientRequest, limit int) ([]model.Patient, error)
	Create(data model.Patient) (*model.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepositroy(db *gorm.DB) PatientRepository {
	return &patientRepository{db}
}

func (r *patientRepository) FindAll(hospitalID string, req dto.SearchPatientRequest, limit int) ([]model.Patient, error) {
	var data []model.Patient

	query := r.db.Where("hospital_id = ?", hospitalID)

	if req.NationalID != nil {
		query = query.Where("national_id = ?", *req.NationalID)
	}

	if req.PassportID != nil {
		query = query.Where("passport_id = ?", *req.PassportID)
	}

	if req.FirstName != nil {
		keyword := "%" + *req.FirstName + "%"
		query = query.Where(
			"(first_name_th ILIKE ? OR first_name_en ILIKE ?)",
			keyword, keyword,
		)
	}

	if req.MiddleName != nil {
		keyword := "%" + *req.MiddleName + "%"
		query = query.Where(
			"(middle_name_th ILIKE ? OR middle_name_en ILIKE ?)",
			keyword, keyword,
		)
	}

	if req.LastName != nil {
		keyword := "%" + *req.LastName + "%"
		query = query.Where(
			"(last_name_th ILIKE ? OR last_name_en ILIKE ?)",
			keyword, keyword,
		)
	}

	if req.DateOfBirth != nil {
		query = query.Where("date_of_birth = ?", *req.DateOfBirth)
	}

	if req.PhoneNumber != nil {
		query = query.Where("phone_number = ?", *req.PhoneNumber)
	}

	if req.Email != nil {
		query = query.Where("email = ?", *req.Email)
	}

	err := query.Limit(limit).Find(&data).Error

	return data, err
}

func (r *patientRepository) Create(data model.Patient) (*model.Patient, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
