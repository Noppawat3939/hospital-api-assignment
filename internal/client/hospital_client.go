package client

import (
	"encoding/json"
	"fmt"
	"hospital-api/internal/dto"
	"net/http"
	"time"
)

type hospitalClient struct {
	baseURL string
	client  *http.Client
}

type HospitalClient interface {
	GetPatientByID(id string) (*dto.HospitalClientPatientResponse, error)
}

func NewHospitalClient(baseURL string) *hospitalClient {
	return &hospitalClient{baseURL: baseURL, client: &http.Client{Timeout: 5 * time.Second}}
}

func (c *hospitalClient) GetPatientByID(id string) (*dto.HospitalClientPatientResponse, error) {
	url := fmt.Sprintf("%s/patient/search/%s", c.baseURL, id)

	res, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HIS api error %d", res.StatusCode)
	}

	var response dto.HospitalClientPatientResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
