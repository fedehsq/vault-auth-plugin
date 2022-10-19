package adminapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	admindao "vault-auth-plugin/server/dao/admin"
	"vault-auth-plugin/server/models/admin"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("secretKey")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Minute).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	JWT      string `json:"jwt"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
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
	token, err := generateJWT()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	admin := Admin{
		Username: user.Username,
		Password: user.Password,
		JWT:      token,
	}
	fmt.Println(token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admin)
}
