package dto

import "time"

type StaffRequestBaseFields struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Hospital string `json:"hospital" binding:"required"`
}

type StaffLoginResult struct {
	AccessToken string `json:"access_token"`
}

type CreateStaffResponse struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	HospitalID string    `json:"hospital"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
