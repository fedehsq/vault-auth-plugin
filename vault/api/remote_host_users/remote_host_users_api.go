package remotehostusersapi

import (
	"encoding/json"
	"fmt"
	"github.com/fedehsq/vault/config"
	"io"
	"net/http"
	"strings"
)

type RemoteHostUserRequest struct {
	RemoteHostIp string `json:"remote_host_ip" example:"192.168.1.1"`
	Username     string `json:"username" example:"fedehsq"`
}

type RemoteHostUsersResponse struct {
	RemoteHostIp string   `json:"remote_host_ip" example:"192.168.1.1"`
	Users        []string `json:"users" example:"[fedehsq, root]"`
}

func (r RemoteHostUsersResponse) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"remote_host_ip": r.RemoteHostIp,
		"users":          r.Users,
	}
}

type RemoteHostUserResponse struct {
	RemoteHostIp string `json:"remote_host_ip" example:"192.168.1.1"`
	Username     string `json:"username" example:"fedehsq"`
}

func (r RemoteHostUserResponse) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"remote_host_ip": r.RemoteHostIp,
		"username":       r.Username,
	}
}

func Create(ip string, username string, jwt string) (*RemoteHostUserResponse, error) {
	rb, err := json.Marshal(RemoteHostUserRequest{
		RemoteHostIp: ip,
		Username:     username,
	})

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/remote-host-users", config.Conf.ApiAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	remoteHostUser := RemoteHostUserResponse{}
	err = json.Unmarshal(body, &remoteHostUser)
	if err != nil {
		return nil, err
	}
	return &remoteHostUser, nil
}

func Get(ip string, username string, jwt string) (*RemoteHostUserResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/remote-host-users?ip=%s&username=%s", config.Conf.ApiAddress, ip, username), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}
	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	remoteHostUser := RemoteHostUserResponse{}
	err = json.Unmarshal(body, &remoteHostUser)
	if err != nil {
		return nil, err
	}
	return &remoteHostUser, nil
}

func GetAll(ip string, jwt string) (*RemoteHostUsersResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/remote-host-users?ip=%s", config.Conf.ApiAddress, ip), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}
	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	remoteHostUsers := RemoteHostUsersResponse{}
	err = json.Unmarshal(body, &remoteHostUsers)
	if err != nil {
		return nil, err
	}
	return &remoteHostUsers, nil
}

func Delete(ip string, username string, jwt string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/remote-host-users?ip=%s&username=%s", config.Conf.ApiAddress, ip, username), strings.NewReader(string("")))

	if err != nil {
		return err
	}
	_, err = doRequest(req, jwt)
	if err != nil {
		return err
	}
	return nil
}

func doRequest(req *http.Request, jwt string) ([]byte, error) {
	req.Header.Set("Authorization", jwt)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}
