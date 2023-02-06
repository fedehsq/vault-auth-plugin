package remotehostapi

import (
	"encoding/json"
	"fmt"
	"github.com/fedehsq/vault/config"
	"io"
	"net/http"
	"strings"
)

type RemoteHostResponse struct {
	Ip string `json:"ip" example:"192.168.1.1"`
}

func Create(ip string, jwt string) (*RemoteHostResponse, error) {
	rb, err := json.Marshal(map[string]string{
		"ip": ip,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/remote-hosts", config.Conf.ApiAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	
	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	remoteHost := RemoteHostResponse{}
	err = json.Unmarshal(body, &remoteHost)
	if err != nil {
		return nil, err
	}
	return &remoteHost, nil
}

func Get(ip string, jwt string) (*RemoteHostResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/remote-hosts?ip=%s", config.Conf.ApiAddress, ip), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}
	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}
	remoteHost := RemoteHostResponse{}
	err = json.Unmarshal(body, &remoteHost)
	if err != nil {
		return nil, err
	}
	return &remoteHost, nil
}

func Delete(ip string, jwt string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/remote-hosts?ip=%s", config.Conf.ApiAddress, ip), strings.NewReader(string("")))
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
