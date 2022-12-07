package logapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fedehsq/api/api"
	"net/http"
	"strings"
	// import main package to get the esClient
	"github.com/fedehsq/api/config"
)

// ListLogs godoc
// @Summary      List logs
// @Description  get all logs
// @Tags         logs
// @Produce      json
// @Success      200  {array}   log.Log
// @Failure      400
// @Failure      403
// @Router       /v1/log/get-logs [get]
// @Security 	 JWT
func GetAll(w http.ResponseWriter, r *http.Request) {
	ok,identity, err := api.VerifyToken(r)
	
	if err != nil {
		api.WriteLog("GET", "/v1/log/get-logs", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("GET", "/v1/log/get-logs", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	api.WriteLog("GET", "/v1/log/get-logs", identity, r.RemoteAddr, "")

	body := fmt.Sprintf(
		`{"query": {"match_all": {}},"sort": [{"time": {"order": "desc"}}]}`)
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
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": logs["hits"],
	})
}
