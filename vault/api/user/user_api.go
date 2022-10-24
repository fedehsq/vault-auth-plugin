package userapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"vault-auth-plugin/vault/api"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT of the bastion host (bastion authenticated itself to the vault before)
func SignIn(username string, password string, jwt string) (*User, error) {
	fmt.Println("signing in")
	rb, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signin", api.VaultServerAddress), strings.NewReader(string(rb)))
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

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", api.VaultServerAddress), strings.NewReader(string(rb)))
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

func GetUsers(jwt string) ([]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", api.VaultServerAddress), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(username string, jwt string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user?username=%s", api.VaultServerAddress, username), strings.NewReader(string("")))
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/user?username=%s", api.VaultServerAddress, username), strings.NewReader(string("")))
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

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/user", api.VaultServerAddress), strings.NewReader(string(rb)))
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
