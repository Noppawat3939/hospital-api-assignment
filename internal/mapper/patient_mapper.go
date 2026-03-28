package mapper

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/model"
	"hospital-api/pkg/common"
	"hospital-api/pkg/timeutil"
	"strings"
	"time"
)

func ToSearchPatientsResponse(patients []model.Patient) []dto.SearchPatientResponse {
	res := make([]dto.SearchPatientResponse, 0, len(patients))

	for _, p := range patients {
		var dob *string

		if p.DateOfBirth != nil {
			formatted := p.DateOfBirth.Format(timeutil.YYYYMMDD)
			dob = &formatted
		}

		res = append(res, dto.SearchPatientResponse{
			FirstNameTH:  p.FirstNameTH,
			MiddleNameTH: p.MiddleNameTH,
			LastNameTH:   p.LastNameTH,
			FirstNameEN:  p.FirstNameEN,
			MiddleNameEN: p.MiddleNameEN,
			LastNameEN:   p.LastNameEN,
			DateOfBirth:  dob,
			PatientHN:    p.PatientHN,
			NationalID:   p.NationalID,
			PassportID:   p.PassportID,
			Email:        p.Email,
			Gender:       string(p.Gender),
		})
	}

	return res
}

func ToCreatePatient(res *dto.HospitalClientPatientResponse, hospitalID string) (model.Patient, error) {
	var dob *time.Time

	if res.DateOfBirth != "" {
		parsed, err := timeutil.ParseDate(res.DateOfBirth)

		if err != nil {
			return model.Patient{}, err
		}

		dob = parsed
	}

	return model.Patient{
		HospitalID:   hospitalID,
		FirstNameTH:  res.FirstNameTH,
		MiddleNameTH: res.MiddleNameTH,
		LastNameTH:   res.LastNameTH,
		FirstNameEN:  res.FirstNameEN,
		MiddleNameEN: res.MiddleNameEN,
		LastNameEN:   res.LastNameEN,
		DateOfBirth:  dob,
		PatientHN:    res.PatientHN,
		NationalID:   common.StringToPtr(res.NationalID),
		PassportID:   common.StringToPtr(res.PassportID),
		PhoneNumber:  res.PhoneNumber,
		Email:        res.Email,
		Gender:       ToGender(res.Gender),
	}, nil
}

func ToGender(v string) model.PatientGender {
	cleaned := strings.ToUpper(strings.TrimSpace(v))
	var g model.PatientGender

	switch cleaned {
	case "M":
		g = model.Male
	case "F":
		g = "F"
	default:
		g = ""
	}

	return g
}
