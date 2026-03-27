package dto

import "time"

type SearchPatientRequest struct {
	NationalID  *string    `json:"national_id"`
	PassportID  *string    `json:"passport_id"`
	FirstName   *string    `json:"first_name"`
	MiddleName  *string    `json:"middle_name"`
	LastName    *string    `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	PhoneNumber *string    `json:"phone_number"`
	Email       *string    `json:"email"`
}
