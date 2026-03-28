package service

import (
	"errors"
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/internal/repository"
	"hospital-api/pkg/jwt"
	"hospital-api/pkg/password"
	"hospital-api/pkg/response"
	"time"
)

type StaffService interface {
	Create(req dto.StaffRequestBaseFields) (*model.Staff, error)
	Login(req dto.StaffRequestBaseFields) (*dto.StaffLoginResult, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo}
}

func (s *staffService) Create(req dto.StaffRequestBaseFields) (*model.Staff, error) {
	hashed := password.Hash(req.Password)

	staff, err := s.repo.Create(model.Staff{Username: req.Username, Password: hashed, HospitalID: req.Hospital})
	if err != nil {
		return nil, err
	}

	return staff, nil
}

func (s *staffService) Login(req dto.StaffRequestBaseFields) (*dto.StaffLoginResult, error) {
	// check not found
	staff, err := s.repo.FindOneByUsernameAndHospitalID(req.Username, req.Hospital)
	if err != nil {
		return nil, errors.New(response.ErrInvalidCrediental)
	}

	// check password
	matched := password.Compare(req.Password, staff.Password)
	if !matched {
		return nil, errors.New(response.ErrInvalidCrediental)
	}

	// gen jwt
	exp := time.Now().Add(1 * time.Hour)

	tk, err := jwt.Gen(staff.Username, staff.HospitalID, exp)
	if err != nil {
		return nil, err
	}

	return &dto.StaffLoginResult{AccessToken: tk}, nil
}
