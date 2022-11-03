package adminapi

import (
	"encoding/json"
	"fmt"
	"github.com/fedehsq/vault/config"
	"io"
	"net/http"
	"strings"
)

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	JWT      string `json:"jwt"`
}

func SignIn(username string, password string) (*Admin, error) {
	rb, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin-signin", config.Conf.ApiAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	admin := Admin{}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func doRequest(req *http.Request) ([]byte, error) {
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
