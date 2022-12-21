package logapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fedehsq/api/api"
	"github.com/fedehsq/api/config"
	"net/http"
	"strings"
)

type LogRequest struct {
	Command string `json:"command"`
	SshAddress string `json:"ssh_address"`
	Username string `json:"username"`
}

// ListLogs godoc
// @Summary      List logs
// @Description  Returns the logs requested; if the parameters are empty returns all
// @Tags         logs
// @Param        value query string false "The ip address of the caller;The identity of the caller; The HTTP method called; The route requested; The command inserted"  Format(string)
// @Produce      json
// @Success      200  {array}   log.Log
// @Failure      400
// @Failure      403
// @Router       /v1/logs/ [get]
// @Security 	 JWT
func Get(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	query := r.URL.Query().Get("q")

	if err != nil {
		api.WriteLog("GET", "/v1/logs", "Unauthorized user", r.RemoteAddr, query)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("GET", "/v1/logs", "Unauthorized user", r.RemoteAddr, query)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	api.WriteLog("GET", "/v1/logs", identity, r.RemoteAddr, query)

	var body string
	if query == "" {
		body =
			`{"query": {"match_all": {}}}`
	} else {
		body = fmt.Sprintf(
			`{"query": {"multi_match": {"query": "%s", "fields": ["ip", "caller_identity", "method", "route", "body", "time"]}}}`,
			query)
	}
	res, err := config.EsClient.Search(
		config.EsClient.Search.WithContext(context.Background()),
		config.EsClient.Search.WithIndex("logs"),
		config.EsClient.Search.WithBody(strings.NewReader(body)),
		config.EsClient.Search.WithPretty(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer res.Body.Close()
	if res.IsError() {
		http.Error(w, res.String(), http.StatusBadRequest)
		return
	}
	var logs map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&logs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode only the hits part of the response body, without the keys as array of strings
	// (the keys are the fields of the struct Log)
	if err := json.NewEncoder(w).Encode(logs["hits"].(map[string]interface{})["hits"]); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Create log godoc
// @Summary      Create log
// @Description  Creates a new log
// @Tags         logs
// @Param        log body log.Log true "The log to create"
// @Produce      json
// @Success      200  {object}   log.Log
// @Failure      400
// @Failure      403
// @Router       /v1/logs/ [post]
// @Security 	 JWT
func Post(w http.ResponseWriter, r *http.Request) {
	ok, _, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("POST", "/v1/logs", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("POST", "/v1/logs", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	var log LogRequest
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.WriteLog("POST", "/v1/logs", log.Username, log.SshAddress, log.Command)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
