package handler

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/middleware"
	"hospital-api/internal/model"
	"hospital-api/internal/service"
	"hospital-api/pkg/response"
	"hospital-api/pkg/timeutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	srv service.PatientService
}

func NewPatientHandler(srv service.PatientService) *PatientHandler {
	return &PatientHandler{srv}
}

func (h *PatientHandler) PatientSearch(c *gin.Context) {
	var req dto.SearchPatientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.ErrBodyInvalid)
		return
	}

	staff, ok := middleware.GetStaffContext(c)

	if !ok {
		response.Error(c, http.StatusUnauthorized, response.ErrUnAuthorized)
		return
	}

	patients, err := h.srv.Search(staff.HospitalID, req)
	if err != nil {
		response.Error(c, http.StatusNotFound, err.Error())
		return
	}

	data := mapPatientsResponse(patients)
	response.Success(c, data)
}

func mapPatientsResponse(patients []model.Patient) []dto.SearchPatientResponse {
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
