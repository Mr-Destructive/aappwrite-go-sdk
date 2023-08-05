package appwrite

import (
	"encoding/json"
	"strings"
)

// Functions service
type Function struct {
	Client Client
}

func NewFunctions(clt Client) Function {
	function := Function{
		Client: clt,
	}

	return function
}

type Variable struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	FunctionId string `json:"functionId"`
	Id         string `json:"$id"`
	CreatedAt  string `json:"$createdAt"`
	UpdatedAt  string `json:"$updatedAt"`
}

type FunctionObject struct {
	Id               string     `json:"$id"`
	Name             string     `json:"name,omitempty"`
	CreatedAt        string     `json:"$createdAt"`
	UpdatedAt        string     `json:"$updatedAt"`
	Execute          []string   `json:"execute"`
	Enabled          bool       `json:"enabled"`
	Variable         []Variable `json:"variable"`
	Runtime          string     `json:"runtime"`
	Deployment       string     `json:"deployment"`
	Events           []string   `json:"events"`
	Schedule         string     `json:"schedule"`
	ScheduleNext     string     `json:"scheduleNext"`
	SchedulePrevious string     `json:"schedulePrevious"`
	Timeout          int        `json:"timeout"`
}

type DeploymentObject struct {
	Id           string `json:"$id"`
	CreatedAt    string `json:"$createdAt"`
	UpdatedAt    string `json:"$updatedAt"`
	ResourceId   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	EntryPoint   string `json:"entryPoint"`
	Size         int    `json:"size"`
	BuildId      string `json:"buildId"`
	Activate     bool   `json:"activate"`
	Status       string `json:"status"`
	BuildStdout  string `json:"buildStdout"`
	BuildStderr  string `json:"buildStderr"`
	BuildTime    string `json:"buildTime"`
}

type ExecutionObject struct {
	Id          string   `json:"$id"`
	CreatedAt   string   `json:"$createdAt"`
	UpdatedAt   string   `json:"$updatedAt"`
	Permissions []string `json:"permissions"`
	FunctionId  string   `json:"functionId"`
	Trigger     string   `json:"trigger"`
	Status      string   `json:"status"`
	StatusCode  int      `json:"statusCode"`
	Response    string   `json:"response"`
	Stdout      string   `json:"stdout"`
	Stderr      string   `json:"stderr"`
	Duration    string   `json:"duration"`
}

type VariableListResponse struct {
	Total     int        `json:"total"`
	Variables []Variable `json:"variables"`
}

type FunctionListResponse struct {
	Total     int              `json:"total"`
	Functions []FunctionObject `json:"functions"`
}

type DeploymentListResponse struct {
	Total       int                `json:"total"`
	Deployments []DeploymentObject `json:"deployments"`
}

type ExecutionListResponse struct {
	Total      int               `json:"total"`
	Executions []ExecutionObject `json:"executions"`
}

func (srv *Function) ListFunctions(Search string, Queries []string) (*FunctionListResponse, error) {
	path := "/functions"
	params := map[string]interface{}{
		"search":  Search,
		"queries": Queries,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result FunctionListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) GetFunction(functionId string) (*FunctionObject, error) {
	r := strings.NewReplacer("{functionId}", functionId)
	path := r.Replace("/functions/{functionId}")
	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result FunctionObject
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) ListDeployments(functionId, Search string, Queries []string) (*DeploymentListResponse, error) {
	r := strings.NewReplacer("{functionId}", functionId)
	path := r.Replace("/functions/{functionId}/deployments")
	params := map[string]interface{}{
		"search":  Search,
		"queries": Queries,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result DeploymentListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) GetDeployment(functionId, deploymentId string) (*DeploymentObject, error) {
	r := strings.NewReplacer("{functionId}", functionId, "{deploymentId}", deploymentId)
	path := r.Replace("/functions/{functionId}/deployments/{deploymentId}")
	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result DeploymentObject
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) ListExecutions(functionId, Search string, Queries []string) (*ExecutionListResponse, error) {
	r := strings.NewReplacer("{functionId}", functionId)
	path := r.Replace("/functions/{functionId}/executions")
	params := map[string]interface{}{
		"search":  Search,
		"queries": Queries,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result ExecutionListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) GetExecution(functionId, executionId string) (*ExecutionObject, error) {
	r := strings.NewReplacer("{functionId}", functionId, "{executionId}", executionId)
	path := r.Replace("/functions/{functionId}/executions/{executionId}")
	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result ExecutionObject
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) ListVariables(functionId, Search string, Queries []string) (*VariableListResponse, error) {
	path := "/functions/{functionId}/variables"
	params := map[string]interface{}{
		"search":  Search,
		"queries": Queries,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result VariableListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (srv *Function) GetVariable(functionId, variableId string) (*Variable, error) {
	r := strings.NewReplacer("{functionId}", functionId, "{variableId}", variableId)
	path := r.Replace("/functions/{functionId}/variables/{variableId}")
	params := map[string]interface{}{}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result Variable
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
