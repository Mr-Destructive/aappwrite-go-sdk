package appwrite

import (
	"encoding/json"
)

type HealthStatus struct {
	Ping   int    `json:"ping"`
	Status string `json:"status"`
}

type HealthQueue struct {
	Size int `json:"size"`
}

type HealthTime struct {
	RealTime  int `json:"realTime"`
	LocalTime int `json:"localTime"`
	Diff      int `json:"diff"`
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

func (srv *Client) CacheHealth() (*HealthStatus, error) {
	path := "/health/cache"
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

func (srv *Client) LocalStorageHealth() (*HealthStatus, error) {
	path := "/health/storage/local"
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

func (srv *Client) FunctionsQueue() (*HealthQueue, error) {
	path := "/health/queue/functions"
	resp, err := srv.CallAPI("GET", path, srv.headers, nil)
	if err != nil {
		return nil, err
	}
	var result HealthQueue
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Client) LogsQueue() (*HealthQueue, error) {
	path := "/health/queue/logs"
	resp, err := srv.CallAPI("GET", path, srv.headers, nil)
	if err != nil {
		return nil, err
	}
	var result HealthQueue
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Client) WebHooksQueue() (*HealthQueue, error) {
	path := "/health/queue/webhooks"
	resp, err := srv.CallAPI("GET", path, srv.headers, nil)
	if err != nil {
		return nil, err
	}
	var result HealthQueue
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Client) TimeHealth() (*HealthTime, error) {
	path := "/health/time"
	resp, err := srv.CallAPI("GET", path, srv.headers, nil)
	if err != nil {
		return nil, err
	}
	var result HealthTime
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
