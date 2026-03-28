package handler

import (
	"hospital-api/internal/dto"
	"hospital-api/internal/service"
	"hospital-api/pkg/db"
	"hospital-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	srv    service.StaffService
	hosSrv service.HospitalService
}

func NewStaffHandler(srv service.StaffService, hosSrv service.HospitalService) *StaffHandler {
	return &StaffHandler{srv: srv, hosSrv: hosSrv}
}

func (h *StaffHandler) StaffCreate(c *gin.Context) {
	var req dto.StaffRequestBaseFields

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.ErrBodyInvalid)
		return
	}

	hospital, err := h.hosSrv.FindOne(req.Hospital)
	if err != nil || hospital == nil {
		response.Error(c, http.StatusBadRequest, response.ErrHospitalInvalid)
		return
	}

	req.Normalize()

	staff, err := h.srv.Create(req)
	if err != nil {
		if db.IsUniqueConstraintError(err) {
			response.Error(c, http.StatusConflict, response.ErrDupCreateStaff)
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	data := dto.CreateStaffResponse{
		ID:         staff.ID,
		Username:   staff.Username,
		HospitalID: staff.HospitalID,
		CreatedAt:  staff.CreatedAt,
		UpdatedAt:  staff.UpdatedAt,
	}
	response.Success(c, data)
}

func (h *StaffHandler) StaffLogin(c *gin.Context) {
	var req dto.StaffRequestBaseFields

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.ErrBodyInvalid)
		return
	}

	result, err := h.srv.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	data := dto.StaffLoginResult{AccessToken: result.AccessToken}

	response.Success(c, data)
}
