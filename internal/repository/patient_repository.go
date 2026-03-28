package repository

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"

	"gorm.io/gorm"
)

const (
	defaultLimit = 20
)

type PatientRepository interface {
	FindAll(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error)
	FindOneByIdentity(req dto.SearchPatientRequest) (*model.Patient, error)
	Create(data model.Patient) (*model.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepositroy(db *gorm.DB) PatientRepository {
	return &patientRepository{db}
}

func (r *patientRepository) FindAll(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
	var data []model.Patient

	// safe limit: ถ้า req.Limit เป็น nil ใช้ default 100
	limit := 100
	// if req.Limit != nil && *req.Limit > 0 {
	// 	limit = *req.Limit
	// }

	query := r.db.Where("hospital_id = ?", hospitalID)

	// Search by national_id OR passport_id
	nationalID := ""
	passportID := ""
	if req.NationalID != nil {
		nationalID = *req.NationalID
	}
	if req.PassportID != nil {
		passportID = *req.PassportID
	}

	if nationalID != "" && passportID != "" {
		query = query.Where("national_id = ? OR passport_id = ?", nationalID, passportID)
	} else if nationalID != "" {
		query = query.Where("national_id = ?", nationalID)
	} else if passportID != "" {
		query = query.Where("passport_id = ?", passportID)
	}

	// Additional filters (AND)
	if req.FirstName != nil && *req.FirstName != "" {
		keyword := "%" + *req.FirstName + "%"
		query = query.Where("(first_name_th ILIKE ? OR first_name_en ILIKE ?)", keyword, keyword)
	}

	if req.MiddleName != nil && *req.MiddleName != "" {
		keyword := "%" + *req.MiddleName + "%"
		query = query.Where("(middle_name_th ILIKE ? OR middle_name_en ILIKE ?)", keyword, keyword)
	}

	if req.LastName != nil && *req.LastName != "" {
		keyword := "%" + *req.LastName + "%"
		query = query.Where("(last_name_th ILIKE ? OR last_name_en ILIKE ?)", keyword, keyword)
	}

	if req.DateOfBirth != nil {
		query = query.Where("date_of_birth = ?", *req.DateOfBirth)
	}

	if req.PhoneNumber != nil && *req.PhoneNumber != "" {
		query = query.Where("phone_number = ?", *req.PhoneNumber)
	}

	if req.Email != nil && *req.Email != "" {
		query = query.Where("email = ?", *req.Email)
	}

	// Limit
	err := query.Limit(limit).Find(&data).Error
	return data, err
}

func (r *patientRepository) FindOneByIdentity(req dto.SearchPatientRequest) (*model.Patient, error) {
	var data model.Patient
	query := r.db

	if req.PassportID != nil && *req.PassportID != "" {
		query = query.Where("passport_id = ?", *req.PassportID)
	}

	if req.NationalID != nil && *req.NationalID != "" {
		query = query.Where("national_id = ?", *req.NationalID)
	}

	err := query.First(&data).Error

	return &data, err
}

func (r *patientRepository) Create(data model.Patient) (*model.Patient, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
