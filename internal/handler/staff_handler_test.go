package handler

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// staff service
type mockStaffService struct {
	CreateFunc func(req dto.StaffRequestBaseFields) (*model.Staff, error)
	LoginFunc  func(req dto.StaffRequestBaseFields) (*dto.StaffLoginResult, error)
}

func (m *mockStaffService) Create(req dto.StaffRequestBaseFields) (*model.Staff, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(req)
	}
	return nil, nil
}

func (m *mockStaffService) Login(req dto.StaffRequestBaseFields) (*dto.StaffLoginResult, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(req)
	}
	return nil, nil
}

// hospital service
type mockHospitalService struct {
	FindOneFunc func(hospitalID string) (*model.Hospital, error)
}

func (m *mockHospitalService) FindOne(hospitalID string) (*model.Hospital, error) {
	if m.FindOneFunc != nil {
		return m.FindOneFunc(hospitalID)
	}
	return nil, nil
}

func TestStaffHandler_StaffCreate_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSrv := &mockStaffService{
		CreateFunc: func(req dto.StaffRequestBaseFields) (*model.Staff, error) {
			return &model.Staff{
				ID:         1,
				Username:   req.Username,
				HospitalID: req.Hospital,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}, nil
		},
	}

	mockHosSrv := &mockHospitalService{
		FindOneFunc: func(hospitalID string) (*model.Hospital, error) {
			return &model.Hospital{ID: hospitalID, Name: "Hospital A"}, nil
		},
	}

	handler := NewStaffHandler(mockSrv, mockHosSrv)

	body := `{"username":"TestUser","password":"pass123","hospital":"hospital_a"}`
	req := httptest.NewRequest(http.MethodPost, "/staff", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	handler.StaffCreate(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"ID":1`)
	assert.Contains(t, w.Body.String(), `"Username":"TestUser"`)
	assert.Contains(t, w.Body.String(), `"HospitalID":"hospital_a"`)
}
