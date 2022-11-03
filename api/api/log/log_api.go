package logapi

import (
	"encoding/json"
	"github.com/fedehsq/api/api"
	"github.com/fedehsq/api/dao/audit"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("GetAll Logs", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	audits, err := auditdao.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(audits)
}
