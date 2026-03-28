package service

import (
	"hospital-api/internal/client"
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/internal/repository"
	"hospital-api/pkg/timeutil"
	"strings"
	"time"

	"gorm.io/gorm"
)

type PatientService interface {
	Search(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error)
}

type patientService struct {
	repo   repository.PatientRepository
	client client.HospitalClient
}

func NewPatientService(repo repository.PatientRepository, client client.HospitalClient) PatientService {
	return &patientService{repo, client}
}

func (s *patientService) Search(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
	patients, err := s.repo.FindAll(hospitalID, req)

	if err != nil {
		return nil, err
	}

	// has patients
	if len(patients) > 0 {
		return patients, nil
	}

	// check identity of own hospital
	existing, err := s.repo.FindOneByIdentity(req)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// found but other hospital
	if existing != nil && existing.HospitalID != "" && existing.HospitalID != hospitalID {
		return []model.Patient{}, nil
	}

	var searchID string = ""

	if req.PassportID != nil && *req.PassportID != "" {
		searchID = *req.PassportID
	} else if req.NationalID != nil && *req.NationalID != "" {
		searchID = *req.NationalID
	}

	// id not found in query
	if searchID == "" {
		return []model.Patient{}, nil
	}

	// fetch from HIS (mock)
	res, err := s.client.GetPatientByID(searchID)

	if err != nil {
		return nil, err
	}

	mapped, err := mapPatientToCreate(res, hospitalID)
	if err != nil {
		return nil, err
	}

	// create new patient
	patient, err := s.repo.Create(mapped)
	if err != nil {
		return nil, err
	}

	return []model.Patient{*patient}, nil
}

func mapPatientToCreate(res *client.PatientResponse, hospitalID string) (model.Patient, error) {
	var dob *time.Time

	if res.DateOfBirth != "" {
		parsed, err := timeutil.ParseData(res.DateOfBirth)

		if err != nil {
			return model.Patient{}, err
		}

		dob = parsed
	}

	return model.Patient{
		HospitalID:   hospitalID,
		FirstNameTH:  res.FirstNameTH,
		MiddleNameTH: res.MiddleNameTH,
		LastNameTH:   res.LastNameTH,
		FirstNameEN:  res.FirstNameEN,
		MiddleNameEN: res.MiddleNameEN,
		LastNameEN:   res.LastNameEN,
		DateOfBirth:  dob,
		PatientHN:    res.PatientHN,
		NationalID:   res.NationalID,
		PassportID:   stringToPtr(res.PassportID),
		PhoneNumber:  res.PhoneNumber,
		Email:        res.Email,
		Gender:       mapGender(res.Gender),
	}, nil
}

func mapGender(v string) model.PatientGender {
	cleaned := strings.ToUpper(strings.TrimSpace(v))

	var data model.PatientGender

	switch cleaned {
	case "M":
		data = model.Male
	case "F":
		data = "F"
	default:
		data = ""
	}

	return data
}

func stringToPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
