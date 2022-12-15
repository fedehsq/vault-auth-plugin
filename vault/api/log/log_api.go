package logapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/fedehsq/vault/config"
)

// Create a struct with this structure:
/* [{
    "_id": "33",
    "_index": "logs",
    "_score": 1,
    "_source": {
        "body": "2022",
        "caller_identity": "Unauthorized user",
        "id": 33,
        "ip": "127.0.0.1:52954",
        "method": "GET",
        "route": "/v1/logs",
        "time": "2022-12-12T10:46:58.116079Z"
    }
}] */
type ElkResponse struct {
	Id     string `json:"_id"`
	Index  string `json:"_index"`
	Score  int    `json:"_score"`
	Source Log    `json:"_source"`
}

func (e *ElkResponse) String() string {
	return e.Source.String()
}

type Log struct {
	Body           string    `json:"body" example:"{username: admin, password: ********}"`
	CallerIdentity string    `json:"caller_identity" example:"admin"`
	Ip             string    `json:"ip" example:"127.0.0.1:50336"`
	Method         string    `json:"method" example:"POST"`
	Route          string    `json:"route" example:"/api/v1/admin/signin"`
	Time           time.Time `json:"time" example:"2022-10-27 10:18:47.791249"`
}

func (l *Log) String() string {
	return fmt.Sprintf("Body: %s, Caller identity: %s, Ip: %s, Method: %s, Route:%s, Time: %s", l.Body, l.CallerIdentity, l.Ip, l.Method, l.Route, l.Time.String())
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
