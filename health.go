package appwrite

import (
	"encoding/json"
)

type HealthStatus struct {
	Ping   int    `json:"ping"`
	Status string `json:"status"`
}

func (srv *Client) Health() (*HealthStatus, error) {
	path := "/health"
	resp, err := srv.CallAPI("GET", path, srv.headers, nil)
	if err != nil {
		return nil, err
	}
	var result HealthStatus
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Client) DBHealth() (*HealthStatus, error) {
	path := "/health/db"
	resp, err := srv.CallAPI("GET", path, srv.headers, nil)
	if err != nil {
		return nil, err
	}
	var result HealthStatus
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}


