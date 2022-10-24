package adminapi

import (
	"encoding/json"
	"net/http"
	"vault-auth-plugin/vault_server/api"
	"vault-auth-plugin/vault_server/dao/admin"
	"vault-auth-plugin/vault_server/models/admin"
)

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	JWT      string `json:"jwt"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Signin Admin", r)
	var p admin.Admin
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := admindao.GetByUsername(p.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user == nil {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	if user.Password != p.Password {
		http.Error(w, "Wrong password", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := api.GenerateJWT()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	admin := Admin{
		Username: user.Username,
		Password: user.Password,
		JWT:      token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}
