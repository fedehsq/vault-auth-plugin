package userapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	bastionhostapi "vault-auth-plugin/bastion_host/api/bastion_host"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Auth struct {
		ClientToken string `json:"client_token"`
	} `json:"auth"`
}

const (
	host = "http://127.0.0.1:8200"
)

func Signin(w http.ResponseWriter, r *http.Request) {

	var p UserRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Bastion host authentication with vault
	bh, err := bastionhostapi.SignIn()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// User authentication with vault: as authorizion header pass the vault token and jwt
	rb, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/auth/auth-plugin/login", host), strings.NewReader(string(rb)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := doRequest(req, bh.Auth.ClientToken, bh.Auth.Jwt.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := UserResponse{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func doRequest(req *http.Request, vaultToken string, JWT string) ([]byte, error) {
	req.Header.Set("X-Vault-Token", vaultToken)
	req.Header.Set("Authorization", JWT)

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
