package logapi

import (
	"encoding/json"
	"fmt"
	"github.com/fedehsq/vault/config"
	"io"
	"net/http"
	"strings"
)

type ElkResponse struct {
	Id     string  `json:"_id"`
	Index  string  `json:"_index"`
	Score  float64 `json:"_score"`
	Source LogResponse     `json:"_source"`
}

func (e *ElkResponse) String() string {
	return e.Source.String()
}

type LogResponse struct {
	Body           string `json:"body" example:"{username: admin, password: ********}"`
	CallerIdentity string `json:"caller_identity" example:"admin"`
	Ip             string `json:"ip" example:"127.0.0.1:50336"`
	Method         string `json:"method" example:"POST"`
	Route          string `json:"route" example:"/api/v1/admin/signin"`
	Time           string `json:"time" example:"Tue Nov 10 23:00:00 UTC 2009"`
}

type LogRequest struct {
	Command string `json:"command"`
	SshAddress string `json:"ssh_address"`
	Username string `json:"username"`
}

func (l *LogResponse) String() string {
	return fmt.Sprintf("Body: %s, Caller identity: %s, Ip: %s, Method: %s, Route:%s, Time: %s", l.Body, l.CallerIdentity, l.Ip, l.Method, l.Route, l.Time)
}

func Get(jwt string, q string) ([]ElkResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/logs?q=%s", config.Conf.ApiAddress, q), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}

	var logs []ElkResponse
	err = json.Unmarshal(body, &logs)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func Create(command string, sshAddress string, username string, jwt string) error {
	log := LogRequest{
		Command: command,
		SshAddress: sshAddress,
		Username: username,
	}
	rb, err := json.Marshal(log)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/logs", config.Conf.ApiAddress), strings.NewReader(string(rb)))
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
