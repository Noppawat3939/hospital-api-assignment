package handler

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/mapper"
	"hospital-api/internal/middleware"
	"hospital-api/internal/service"
	"hospital-api/pkg/response"
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

	data := mapper.ToSearchPatientsResponse(patients)
	response.Success(c, data)
}
