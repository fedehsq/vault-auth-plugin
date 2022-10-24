package bastionhostapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Token struct {
	Auth struct {
		ClientToken string `json:"client_token"`
		Jwt         struct {
			Token string `json:"jwt"`
		} `json:"metadata"`
	} `json:"auth"`
}

func (t *Token) String() string {
	return fmt.Sprintf("client_token: %s, jwt: %s", t.Auth.ClientToken, t.Auth.Jwt.Token)
}

const (
	host = "http://127.0.0.1:8200"
)

func SignIn() (*Token, error) {
	rb, err := json.Marshal(map[string]string{
		"username": "admin",
		"password": "admin",
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/auth/auth-plugin/admin-login", host), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	tokens := Token{}
	err = json.Unmarshal(body, &tokens)
	if err != nil {
		return nil, err
	}
	return &tokens, nil
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
