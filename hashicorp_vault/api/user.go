package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// ReqUser -
type ReqUser struct {
	HTTPClient *http.Client
	HostURL    string
	Username   string
	Password   string
}

type RespUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	host = "http://localhost:19090"
)

func (c *ReqUser) SignIn() (int, error) {
	rb, err := json.Marshal(map[string]string{
		"username": c.Username,
		"password": c.Password,
	})
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signin", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return -1, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return -1, err
	}
	if body == nil {
		return -1, errors.New("No body")
	}
	return 1, nil
}

func getUsers() ([]RespUser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", host), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var users []RespUser
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUser(username string) (*ReqUser, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", host, username), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err

}

func deleteUser(username string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%s", host, username), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %d", res.StatusCode)
	}

	return nil
}

func (c *ReqUser) doRequest(req *http.Request) ([]byte, error) {

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func CreateUser(username string, password string) (*ReqUser, error) {
	c := &ReqUser{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL:  host,
		Username: username,
		Password: password,
	}

	status, err := c.SignIn()
	if err != nil {
		return nil, err
	}
	if status != 1 {
		return nil, errors.New("Could not sign in")
	}

	if err != nil {
		return nil, err
	}
	return c, nil
}
