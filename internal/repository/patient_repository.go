package repository

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/pkg/pagination"

	"gorm.io/gorm"
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

	query := r.db.Where("hospital_id = ?", hospitalID)

	// exact match filters
	exactFilters := map[string]*string{
		"national_id":  req.NationalID,
		"passport_id":  req.PassportID,
		"phone_number": req.PhoneNumber,
		"email":        req.Email,
	}

	for col, val := range exactFilters {
		if val != nil && *val != "" {
			query = query.Where(col+" = ?", *val)
		}
	}

	// ILIKE filters
	ilikeFilters := map[string]*string{
		"first_name":  req.FirstName,
		"middle_name": req.MiddleName,
		"last_name":   req.LastName,
	}

	for field, val := range ilikeFilters {
		if val != nil && *val != "" {
			kw := "%" + *val + "%"
			query = query.Where("("+field+"_th ILIKE ? OR "+field+"_en ILIKE ?)", kw, kw)
		}
	}

	if req.DateOfBirth != nil {
		query = query.Where("date_of_birth = ?", *req.DateOfBirth)
	}

	p := pagination.Pagination{Limit: req.Limit, Offset: req.Page}
	query = pagination.Apply(query, p)

	err := query.Find(&data).Error
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
