package userapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fedehsq/vault/config"
	"io"
	"net/http"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT of the bastion host (bastion authenticated itself to the vault before)
func SignIn(username string, password string, jwt string) (*User, error) {
	rb, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/users/signin", config.Conf.ApiAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func SignUp(username string, password string, jwt string) (*User, error) {
	rb, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/users", config.Conf.ApiAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(username string, jwt string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/users?username=%s", config.Conf.ApiAddress, username), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}
	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(username string, jwt string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/users?username=%s", config.Conf.ApiAddress, username), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return err
	}

	if string(body) != "DELETED" {
		return errors.New(string(body))
	}

	return nil
}

func UpdateUser(username string, password string, jwt string) (*User, error) {
	rb, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/users", config.Conf.ApiAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
