package service

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock patient repository
type mockPatientRepository struct {
	FindAllFunc           func(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error)
	FindOneByIdentityFunc func(req dto.SearchPatientRequest) (*model.Patient, error)
	CreateFunc            func(data model.Patient) (*model.Patient, error)
}

func (m *mockPatientRepository) FindAll(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc(hospitalID, req)
	}
	return nil, nil
}

func (m *mockPatientRepository) FindOneByIdentity(req dto.SearchPatientRequest) (*model.Patient, error) {
	if m.FindOneByIdentityFunc != nil {
		return m.FindOneByIdentityFunc(req)
	}
	return nil, nil
}

func (m *mockPatientRepository) Create(data model.Patient) (*model.Patient, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(data)
	}
	return nil, nil
}

// mock client
type mockHospitalClient struct {
	GetPatientByIDFunc func(id string) (*dto.HospitalClientPatientResponse, error)
}

func (m *mockHospitalClient) GetPatientByID(id string) (*dto.HospitalClientPatientResponse, error) {
	if m.GetPatientByIDFunc != nil {
		return m.GetPatientByIDFunc(id)
	}
	return nil, nil
}

func TestPatientService_Search_CurrentHospital(t *testing.T) {
	mockRepo := &mockPatientRepository{
		FindAllFunc: func(hospitalID string, req dto.SearchPatientRequest) ([]model.Patient, error) {
			return []model.Patient{
				{PatientHN: "HN1", HospitalID: hospitalID},
			}, nil
		},
	}

	service := NewPatientService(mockRepo, &mockHospitalClient{})

	req := dto.SearchPatientRequest{}
	results, err := service.Search("HOSP1", req)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "HN1", results[0].PatientHN)
}
