package service

import (
	"errors"
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/pkg/password"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStaffRepository struct {
	CreateFunc                         func(data model.Staff) (*model.Staff, error)
	FindOneByUsernameAndHospitalIDFunc func(username, hospitalID string) (*model.Staff, error)
}

func (m *mockStaffRepository) Create(data model.Staff) (*model.Staff, error) {
	return m.CreateFunc(data)
}

func (m *mockStaffRepository) FindOneByUsernameAndHospitalID(username, hospitalID string) (*model.Staff, error) {
	return m.FindOneByUsernameAndHospitalIDFunc(username, hospitalID)
}

func TestStaffService_Create(t *testing.T) {
	mockRepo := &mockStaffRepository{
		CreateFunc: func(data model.Staff) (*model.Staff, error) {
			data.Password = "hashed_" + data.Password
			return &data, nil
		},
	}

	svc := NewStaffService(mockRepo)

	req := dto.StaffRequestBaseFields{
		Username: "testuser",
		Password: "password123",
		Hospital: "hospital_a",
	}

	staff, err := svc.Create(req)

	assert.NoError(t, err)
	assert.Equal(t, "testuser", staff.Username)
	assert.Equal(t, "hospital_a", staff.HospitalID)
	assert.Contains(t, staff.Password, "hashed_")
}

func TestStaffService_Login_Success(t *testing.T) {
	mockRepo := &mockStaffRepository{
		FindOneByUsernameAndHospitalIDFunc: func(username, hospitalID string) (*model.Staff, error) {
			return &model.Staff{
				Username:   username,
				Password:   password.Hash("password123"),
				HospitalID: hospitalID,
			}, nil
		},
	}

	svc := NewStaffService(mockRepo)

	req := dto.StaffRequestBaseFields{
		Username: "testuser",
		Password: "password123",
		Hospital: "hospital_a",
	}

	res, err := svc.Login(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, res.AccessToken)
}

func TestStaffService_Login_InvalidPassword(t *testing.T) {
	mockRepo := &mockStaffRepository{
		FindOneByUsernameAndHospitalIDFunc: func(username, hospitalID string) (*model.Staff, error) {
			return &model.Staff{
				Username:   username,
				Password:   password.Hash("password123"),
				HospitalID: hospitalID,
			}, nil
		},
	}

	svc := NewStaffService(mockRepo)

	req := dto.StaffRequestBaseFields{
		Username: "testuser",
		Password: "wrongpassword",
		Hospital: "hospital_a",
	}

	res, err := svc.Login(req)

	assert.Nil(t, res)
	assert.Error(t, err)
}

func TestStaffService_Login_UserNotFound(t *testing.T) {
	mockRepo := &mockStaffRepository{
		FindOneByUsernameAndHospitalIDFunc: func(username, hospitalID string) (*model.Staff, error) {
			return nil, errors.New("not found")
		},
	}

	svc := NewStaffService(mockRepo)

	req := dto.StaffRequestBaseFields{
		Username: "unknown",
		Password: "password123",
		Hospital: "hospital_a",
	}

	res, err := svc.Login(req)

	assert.Nil(t, res)
	assert.Error(t, err)
}
