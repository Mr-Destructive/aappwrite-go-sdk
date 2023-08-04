package appwrite

import (
	"encoding/json"
	"strings"
)

// Users service
type Users struct {
	Client Client
}

type UserObject struct {
	Id                string                 `json:"$id"`
	Name              string                 `json:"name,omitempty"`
	Email             string                 `json:"email,omitempty"`
	Status            bool                   `json:"status,omitempty"`
	Phone             string                 `json:"phone,omitempty"`
	CreatedAt         string                 `json:"$createdAt"`
	UpdatedAt         string                 `json:"$updatedAt"`
	EmailVerification bool                   `json:"emailVerification"`
	PhoneVerification bool                   `json:"phoneVerification"`
	Registration      string                 `json:"registration"`
	Prefs             map[string]interface{} `json:"prefs"`
	PasswordUpdatedAt string                 `json:"passwordUpdated"`
	Password          string                 `json:"password"`
	HashOptions       map[string]int         `json:"hashOptions"`
	Hash              string                 `json:"hash"`
}

type UserListResponse struct {
	Sum    int          `json:"sum"`
	Offset int          `json:"offset"`
	Limit  int          `json:"limit"`
	Users  []UserObject `json:"users"`
}

func NewUsers(clt Client) Users {
	service := Users{
		Client: clt,
	}

	return service
}

// List get a list of all the project users. You can use the query params to
// filter your results.
func (srv *Users) List(Search string, Limit int, Offset int, OrderType string) ([]UserObject, error) {
	path := "/users"

	params := map[string]interface{}{
		"search":    Search,
		"limit":     Limit,
		"offset":    Offset,
		"orderType": OrderType,
	}

	resp, err := srv.Client.CallAPI("GET", path, srv.Client.headers, params)
	if err != nil {
		return nil, err
	}
	var result UserListResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result.Users, nil
}

// Create create a new user.
func (srv *Users) Create(Email string, Password string, Name string) (map[string]interface{}, error) {
	path := "/users"

	params := map[string]interface{}{
		"email":    Email,
		"password": Password,
		"name":     Name,
	}

	return srv.Client.Call("POST", path, srv.Client.headers, params)
}

// Get get user by its unique ID.
func (srv *Users) Get(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{}

	return srv.Client.Call("GET", path, srv.Client.headers, params)
}

// GetLogs get user activity logs list by its unique ID.
func (srv *Users) GetLogs(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")

	params := map[string]interface{}{}

	return srv.Client.Call("GET", path, srv.Client.headers, params)
}

// GetPrefs get user preferences by its unique ID.
func (srv *Users) GetPrefs(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{}

	return srv.Client.Call("GET", path, srv.Client.headers, params)
}

// UpdatePrefs update user preferences by its unique ID. You can pass only the
// specific settings you wish to update.
func (srv *Users) UpdatePrefs(UserId string, Prefs map[string]interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{
		"prefs": Prefs,
	}

	return srv.Client.Call("PATCH", path, srv.Client.headers, params)
}

// GetSessions get user sessions list by its unique ID.
func (srv *Users) GetSessions(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{}

	return srv.Client.Call("GET", path, srv.Client.headers, params)
}

// DeleteSessions delete all user sessions by its unique ID.
func (srv *Users) DeleteSessions(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{}

	return srv.Client.Call("DELETE", path, srv.Client.headers, params)
}

// DeleteSession delete user sessions by its unique ID.
func (srv *Users) DeleteSession(UserId string, SessionId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions/:session")

	params := map[string]interface{}{
		"sessionId": SessionId,
	}

	return srv.Client.Call("DELETE", path, srv.Client.headers, params)
}

// UpdateStatus update user status by its unique ID.
func (srv *Users) UpdateStatus(UserId string, Status string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/status")

	params := map[string]interface{}{
		"status": Status,
	}

	return srv.Client.Call("PATCH", path, srv.Client.headers, params)
}
