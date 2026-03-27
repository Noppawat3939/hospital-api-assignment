package service

import (
	"fmt"
	"hospital-api/internal/client"
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/internal/repository"
	"time"
)

type PatientService interface {
	Search(hospitalID string, req dto.SearchPatientRequest, limit int) ([]model.Patient, error)
}

type patientService struct {
	repo   repository.PatientRepository
	client client.HospitalClient
}

func NewPatientService(repo repository.PatientRepository, client client.HospitalClient) PatientService {
	return &patientService{repo, client}
}

func (s *patientService) Search(hospitalID string, req dto.SearchPatientRequest, limit int) ([]model.Patient, error) {
	patients, err := s.repo.FindAll(hospitalID, req, limit)
	if err != nil {
		return nil, err
	}

	// has patient
	if len(patients) > 0 {
		return patients, nil
	}

	var searchID string = ""

	if req.PassportID != nil && *req.PassportID != "" {
		searchID = *req.PassportID
	} else if req.NationalID != nil && *req.NationalID != "" {
		searchID = *req.NationalID
	}

	// id not found in query
	if searchID == "" {
		return []model.Patient{}, err
	}
	fmt.Println("searchID", searchID, "hosID", hospitalID)
	// fetch from HIS (mock)
	res, err := s.client.GetPatientByID(searchID, hospitalID)
	if err != nil {
		return nil, err
	}

	patient, err := mapToPatientModel(res, hospitalID)
	if err != nil {
		return nil, err
	}

	// create new patient
	data, err := s.repo.Create(patient)
	if err != nil {
		return nil, err
	}

	return []model.Patient{*data}, nil
}

func mapToPatientModel(res *client.PatientResponse, hospitalID string) (model.Patient, error) {
	var date_of_birth *time.Time

	if res.DateOfBirth != "" {
		t, err := time.Parse("2006-01-02", res.DateOfBirth)
		if err == nil {
			date_of_birth = &t
		}
	}

	return model.Patient{
		HospitalID:   hospitalID,
		FirstNameTH:  res.FirstNameTH,
		MiddleNameTH: res.MiddleNameTH,
		LastNameTH:   res.LastNameTH,
		FirstNameEN:  res.FirstNameEN,
		MiddleNameEN: res.MiddleNameEN,
		LastNameEN:   res.LastNameEN,
		DateOfBirth:  date_of_birth,
		PatientHN:    res.PatientHN,
		NationalID:   res.NationalID,
		PassportID:   stringToPtr(res.PassportID),
		PhoneNumber:  res.PhoneNumber,
		Email:        res.Email,
		Gender:       mapGender(res.Gender),
	}, nil
}

func mapGender(v string) model.PatientGender {
	var data model.PatientGender

	if v == string(model.Male) {
		data = "M"
	} else if v == string(model.Femail) {
		data = "F"
	} else {
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
