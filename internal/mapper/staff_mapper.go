package mapper

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
)

func ToCreateStaffResponse(s *model.Staff) dto.CreateStaffResponse {
	return dto.CreateStaffResponse{
		ID:         s.ID,
		Username:   s.Username,
		HospitalID: s.HospitalID,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
	}
}

func ToStaffLoginResult(r *dto.StaffLoginResult) dto.StaffLoginResult {
	return dto.StaffLoginResult{AccessToken: r.AccessToken}
}
