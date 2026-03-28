package service

import (
	"errors"
	"hospital-api/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock repository
type mockHospitalRepository struct {
	FindOneByIDFunc func(id string) (*model.Hospital, error)
}

func (m *mockHospitalRepository) FindOneByID(id string) (*model.Hospital, error) {
	if m.FindOneByIDFunc != nil {
		return m.FindOneByIDFunc(id)
	}
	return nil, nil
}

func TestHospitalService_FindOne(t *testing.T) {
	mockRepo := &mockHospitalRepository{
		FindOneByIDFunc: func(id string) (*model.Hospital, error) {
			if id == "HOSP1" {
				return &model.Hospital{ID: "HOSP1", Name: "Hospital A"}, nil
			}
			return nil, errors.New("not found")
		},
	}

	svc := NewHospitalService(mockRepo)

	// positive case
	h, err := svc.FindOne("HOSP1")
	assert.NoError(t, err)
	assert.NotNil(t, h)
	assert.Equal(t, "Hospital A", h.Name)

	// negative case
	h, err = svc.FindOne("HOSP2")
	assert.Error(t, err)
	assert.Nil(t, h)
}
