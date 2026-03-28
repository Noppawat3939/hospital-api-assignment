package model

import "time"

type PatientGender string

const (
	Male   PatientGender = "M"
	Femail PatientGender = "F"
)

type Patient struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	HospitalID   string        `gorm:"type:uuid;not null;index" json:"hospital_id"`
	FirstNameTH  string        `gorm:"not null" json:"first_name_th"`
	MiddleNameTH *string       `json:"middle_name_th"`
	LastNameTH   string        `gorm:"not null" json:"last_name_th"`
	FirstNameEN  string        `gorm:"not null" json:"first_name_en"`
	MiddleNameEN *string       `json:"middle_name_en"`
	LastNameEN   string        `gorm:"not null" json:"last_name_en"`
	DateOfBirth  *time.Time    `json:"date_of_birth"`
	PatientHN    string        `gorm:"unique;not null" json:"patient_hn"`
	NationalID   *string       `gorm:"unique" json:"national_id"`
	PassportID   *string       `gorm:"unique" json:"passport_id"`
	PhoneNumber  *string       `json:"phone_number"`
	Email        *string       `json:"email"`
	Gender       PatientGender `gorm:"type:patient_gender" json:"gender"`
}
