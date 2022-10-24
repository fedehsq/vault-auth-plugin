package logapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"vault-auth-plugin/config"
)

type Log struct {
	Time    time.Time `json:"time"`
	Ip      string    `json:"ip"`
	Command string    `json:"command"`
}

func (l *Log) String() string {
	return fmt.Sprintf("Timestamp: %s Ip: %s Command: %s", l.Time.Format(time.RFC3339), l.Ip, l.Command)
}

func GetAll(jwt string) ([]Log, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/logs", config.Conf.VaultServerAddress), strings.NewReader(string("")))
	if err != nil {
		return nil, err
	}

	body, err := doRequest(req, jwt)
	if err != nil {
		return nil, err
	}

	var logs []Log
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
