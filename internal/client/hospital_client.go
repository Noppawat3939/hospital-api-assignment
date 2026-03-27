package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HospitalClient struct {
	baseURL string
	client  *http.Client
}

type PatientResponse struct {
	FirstNameTH  string  `json:"first_name_th"`
	MiddleNameTH *string `json:"middle_name_th"`
	LastNameTH   string  `json:"last_name_th"`
	FirstNameEN  string  `json:"first_name_en"`
	MiddleNameEN *string `json:"middle_name_en"`
	LastNameEN   string  `json:"last_name_en"`
	DateOfBirth  string  `json:"date_of_birth"`
	PatientHN    string  `json:"patient_hn"`
	NationalID   string  `json:"national_id"`
	PassportID   string  `json:"passport_id"`
	PhoneNumber  *string `json:"phone_number"`
	Email        *string `json:"email"`
	Gender       string  `json:"gender"`
}

func NewHospitalClient(baseURL string) *HospitalClient {
	return &HospitalClient{baseURL: baseURL, client: &http.Client{Timeout: 5 * time.Second}}
}

func (c *HospitalClient) GetPatientByID(id, hospitalID string) (*PatientResponse, error) {
	url := fmt.Sprintf("%s/hospital/%s/patient/search/%s", c.baseURL, hospitalID, id)

	res, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HIS api error %d", res.StatusCode)
	}

	var response PatientResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
