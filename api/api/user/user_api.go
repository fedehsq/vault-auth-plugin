package userapi

import (
	"encoding/json"
	"fmt"
	"github.com/fedehsq/api/api"
	"github.com/fedehsq/api/dao/user"
	"github.com/fedehsq/api/models/user"
	"net/http"
)

type UserReq struct {
	Username string `json:"username" example:"user"`
	Password string `json:"password" example:"password"`
}

type UserResp struct {
	Username string `json:"username" example:"user"`
	Password string `json:"password" example:"password"`
}

// SignupUser godoc
//
//		@Summary      Signup an user
//		@Description  Signup an user passing username and password in json
//		@Tags         users
//		@Accept       json
//		@Produce      json
//		@Param        account  body      UserReq  true  "Add user"
//		@Success      201      {object}  UserResp
//		@Failure      400
//		@Failure      401
//		@Router       /signup [post]
//	 	@Security 	 JWT
func Signup(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Signup User", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var p user.User
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, _ := userdao.GetByUsername(p.Username)
	if user != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	err = userdao.Insert(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := UserResp{
		Username: p.Username,
		Password: p.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// SigninUser godoc
//
//	@Summary      Signin an user
//	@Description  Signin an user passing username and password in json
//	@Tags         users
//	@Accept       json
//	@Produce      json
//	@Param        user  body      UserReq  true  "Signin user"
//	@Success      200   {object}  UserResp
//	@Failure      400
//	@Failure      401
//	@Failure      404
//	@Router       /signin [post]
//	@Security 	 JWT
func Signin(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Signin User", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var p user.User
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := userdao.GetByUsername(p.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user == nil {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}
	if user.Password != p.Password {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}
	u := UserResp{
		Username: user.Username,
		Password: user.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// DeleteUser godoc
//
//		@Summary      Delete an user
//		@Description  Delete user passing username
//		@Tags         users
//		@Param        username query string false "user to search by username"  Format(string)
//		@Success      200       "DELETED"
//		@Failure      400
//		@Failure      401
//		@Router       /user [delete]
//	 	@Security 	 JWT
func Delete(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Delete User", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = userdao.Delete(r.URL.Query().Get("username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `DELETED`)
}

// UpdateUser godoc
//
//		@Summary      Update an user
//		@Description  Update an user passing username and password in json
//		@Tags         users
//		@Accept       json
//		@Produce      json
//		@Param        account  body      UserReq  true  "Update user"
//		@Success      200      {object}  UserResp
//		@Failure      400
//		@Failure      401
//		@Failure      404
//		@Router       /user [put]
//	 	@Security 	 JWT
func Update(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Update User", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var p user.User
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := userdao.Update(p.Username, p.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	u := UserResp{
		Username: user.Username,
		Password: user.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// GetUsers godoc
//
//		@Summary      Get all users
//		@Description  Get all users
//		@Tags         users
//		@Produce      json
//		@Success      200      {array}  UserResp
//		@Failure      400
//		@Failure      401
//		@Router       /users [get]
//	 	@Security 	 JWT
func GetAll(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Get Users", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	users, err := userdao.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var usersResp []UserResp
	for _, u := range users {
		usersResp = append(usersResp, UserResp{
			Username: u.Username,
			Password: u.Password,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usersResp)
}

// GetUser godoc
//
//		@Summary      Get an user
//		@Description  Get user passing username
//		@Tags         users
//		@Param        username query string false "user to search by username"  Format(string)
//		@Success      200      {object}  UserResp
//		@Failure      400
//		@Failure      401
//		@Failure      404
//		@Router       /user [get]
//	 	@Security 	 JWT
func GetByUsername(w http.ResponseWriter, r *http.Request) {
	api.WriteLog("Get User", r)
	ok, err := api.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := userdao.GetByUsername(r.URL.Query().Get("username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	u := UserResp{
		Username: user.Username,
		Password: user.Password,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
