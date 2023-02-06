package remotehostsapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fedehsq/api/api"
	"github.com/fedehsq/api/dao/remote-host"
	remotehost "github.com/fedehsq/api/models/remote-host"
)

type RemoteHostReq struct {
	Ip string `json:"ip"`
}

type RemoteHostResp struct {
	Ip string `json:"ip"`
}

// CreateRemoteHost godoc
//
//		@Summary      Create a remote host
//		@Description  Create a remote host passing ip in json
//		@Tags         remote-hosts
//		@Accept       json
//		@Produce      json
//		@Param        remote-host  body      RemoteHostReq  true  "Add remote host"
//		@Success      201      {object}  RemoteHostResp
//		@Failure      400
//		@Failure      401
//		@Router       /v1/remote-hosts [post]
//	 	@Security 	 JWT
func CreateRemoteHost(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("POST", "/v1/remote-hosts", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("POST", "/v1/remote-hosts", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var p remotehost.RemoteHost
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.WriteLog("POST", "/v1/remote-hosts", identity, r.RemoteAddr, p.Json())
	user, _ := remotehostdao.GetByIp(p.Ip)
	if user != nil {
		http.Error(w, "Remote host already exists", http.StatusBadRequest)
		return
	}
	err = remotehostdao.Insert(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h := RemoteHostResp{Ip: p.Ip}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h)
}

// DeleteRemoteHost godoc
//
//		@Summary      Delete a remote host
//		@Description  Delete a remote host passing ip in query
//		@Tags         users
//		@Param        ip  query  string  true  "Delete remote host"
//		@Success      200       "DELETED"
//		@Failure      400
//		@Failure      401
//		@Router       /v1/remote_hosts [delete]
//	 	@Security 	 JWT
func Delete(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("DELETE", "/v1/remote_hosts", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("DELETE", "/v1/remote_hosts", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	api.WriteLog("DELETE", "/v1/remote_hosts", identity, r.RemoteAddr, r.URL.Query().Get("ip"))
	err = remotehostdao.Delete(r.URL.Query().Get("ip"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `DELETED`)
}

// GetRemoteHost godoc
//
//		@Summary      Get a remote host
//		@Description  Get a remote host passing ip in query
//		@Tags         users
//		@Param        ip query string false "user to search by ip"  Format(string)
//		@Success      200      {object}  RemoteHostResp
//		@Failure      400
//		@Failure      401
//		@Failure      404
//		@Router       /v1/remote_hosts [get]
//	 	@Security 	 JWT
func GetByIp(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("GET", "/v1/remote_hosts", "Unauthorized user", r.RemoteAddr, r.URL.Query().Get("ip"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("GET", "/v1/remote_hosts", "Unauthorized user", r.RemoteAddr, r.URL.Query().Get("ip"))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	api.WriteLog("GET", "/v1/remote_hosts", identity, r.RemoteAddr, r.URL.Query().Get("ip"))
	host, err := remotehostdao.GetByIp(r.URL.Query().Get("ip"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	rh := RemoteHostResp{Ip: host.Ip}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rh)
}
