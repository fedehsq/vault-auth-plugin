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
	Source Log     `json:"_source"`
}

func (e *ElkResponse) String() string {
	return e.Source.String()
}

type Log struct {
	Body           string `json:"body" example:"{username: admin, password: ********}"`
	CallerIdentity string `json:"caller_identity" example:"admin"`
	Ip             string `json:"ip" example:"127.0.0.1:50336"`
	Method         string `json:"method" example:"POST"`
	Route          string `json:"route" example:"/api/v1/admin/signin"`
	Time           string `json:"time" example:"Tue Nov 10 23:00:00 UTC 2009"`
}

func (l *Log) String() string {
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
