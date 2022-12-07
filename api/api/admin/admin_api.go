package adminapi

import (
	"encoding/json"
	"github.com/fedehsq/api/api"
	"github.com/fedehsq/api/dao/admin"
	"github.com/fedehsq/api/models/admin"
	"net/http"
)

type AdminReq struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password"`
}

type AdminResp struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password"`
	JWT      string `json:"jwt" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

// SigninAdmin godoc
//
//	@Summary      Signin an admin
//	@Description  Signin an admin passing username and password in json
//	@Tags         admin
//	@Accept       json
//	@Produce      json
//	@Param        admin  body      AdminReq  true  "Signin admin"
//	@Success      200      {object}  AdminResp
//	@Failure      400
//	@Failure      401
//	@Failure      404
//	@Router       /v1/admin/signin [post]
func Signin(w http.ResponseWriter, r *http.Request) {
	var p admin.Admin
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.WriteLog("POST", r.URL.Path, "", r.RemoteAddr, p.Username)
	user, err := admindao.GetByUsername(p.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user == nil {
		http.Error(w, "Admin does not exist", http.StatusNotFound)
		return
	}
	if user.Password != p.Password {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}
	token, err := api.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	admin := AdminResp{
		Username: user.Username,
		Password: user.Password,
		JWT:      token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}
