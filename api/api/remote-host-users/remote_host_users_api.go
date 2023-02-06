package remotehostusersapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fedehsq/api/api"
	"github.com/fedehsq/api/dao/remote-host"
	"github.com/fedehsq/api/dao/remote-host-users"
	"github.com/fedehsq/api/dao/user"
)

type RemoteHostUserReq struct {
	RemoteHostIp string `json:"remote_host_ip"`
	Username     string `json:"username"`
}

func (r *RemoteHostUserReq) Json() string {
	return "{\"remote_host_ip\":\"" + r.RemoteHostIp + "\",\"username\":\"" + r.Username + "\"}"
}

type RemoteHostUserResp struct {
	RemoteHostIp string `json:"remote_host_ip"`
	Username     string `json:"username"`
}

type RemoteHostUsersResp struct {
	RemoteHostIp string   `json:"remote_host_ip"`
	Users        []string `json:"users" example:"[fedehsq, root]"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("POST", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("POST", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqBody RemoteHostUserReq
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.WriteLog("POST", "/v1/remote-host-users", identity, r.RemoteAddr, reqBody.Json())
	// get the user by username
	user, _ := userdao.GetByUsername(reqBody.Username)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if user == nil {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	// get the remote host by ip
	remoteHost, _ := remotehostdao.GetByIp(reqBody.RemoteHostIp)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if remoteHost == nil {
		http.Error(w, "Remote host does not exist", http.StatusBadRequest)
		return
	}
	// check if the user is already associated with the remote host
	remoteHostUser, _ := remotehostusersdao.Get(remoteHost.Id, user.Id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if remoteHostUser != nil {
		http.Error(w, "User already associated", http.StatusBadRequest)
		return
	}
	err = remotehostusersdao.Insert(remoteHost.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h := RemoteHostUserResp{RemoteHostIp: remoteHost.Ip, Username: user.Username}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(h)
}

func GetOrGetAll(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("username") == "" {
		GetAll(w, r)
	} else {
		Get(w, r)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("GET", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, fmt.Sprintf("ip=%s, username=%s", r.URL.Query().Get("ip"), r.URL.Query().Get("username")))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("GET", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, fmt.Sprintf("ip=%s, username=%s", r.URL.Query().Get("ip"), r.URL.Query().Get("username")))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	api.WriteLog("GET", "/v1/remote-host-users", identity, r.RemoteAddr, fmt.Sprintf("ip=%s, username=%s", r.URL.Query().Get("ip"), r.URL.Query().Get("username")))

	// get the user by username
	user, _ := userdao.GetByUsername(r.URL.Query().Get("username"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if user == nil {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	// get the remote host by ip
	remoteHost, _ := remotehostdao.GetByIp(r.URL.Query().Get("ip"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if remoteHost == nil {
		http.Error(w, "Remote host does not exist", http.StatusBadRequest)
		return
	}
	remoteHostUser, _ := remotehostusersdao.Get(remoteHost.Id, user.Id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if remoteHostUser == nil {
		http.Error(w, "User not associated", http.StatusBadRequest)
		return
	}
	resp := RemoteHostUserResp{RemoteHostIp: remoteHost.Ip, Username: user.Username}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("GET", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("GET", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	api.WriteLog("GET", "/v1/remote-host-users", identity, r.RemoteAddr, "")

	// get all remote host users for a specific host
	host, _ := remotehostdao.GetByIp(r.URL.Query().Get("ip"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if host == nil {
		http.Error(w, "Remote host does not exist", http.StatusBadRequest)
		return
	}
	remoteHostUsers, err := remotehostusersdao.GetAll(host.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make([]string, len(remoteHostUsers))
	for i, remoteHostUser := range remoteHostUsers {
		user, err := userdao.Get(remoteHostUser.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp[i] = user.Username
	}
	ipUsers := RemoteHostUsersResp{RemoteHostIp: host.Ip, Users: resp}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ipUsers)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	ok, identity, err := api.VerifyToken(r)
	if err != nil {
		api.WriteLog("DELETE", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		api.WriteLog("DELETE", "/v1/remote-host-users", "Unauthorized user", r.RemoteAddr, "")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	api.WriteLog("DELETE", "/v1/remote-host-users", identity, r.RemoteAddr,
		fmt.Sprintf("ip=%s, username=%s", r.URL.Query().Get("ip"), r.URL.Query().Get("username")))

	// get the user by username
	user, _ := userdao.GetByUsername(r.URL.Query().Get("username"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if user == nil {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}
	// get the remote host by ip
	remoteHost, _ := remotehostdao.GetByIp(r.URL.Query().Get("ip"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	if remoteHost == nil {
		http.Error(w, "Remote host does not exist", http.StatusBadRequest)
		return
	}
	err = remotehostusersdao.Delete(remoteHost.Id, user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `DELETED`)
}
