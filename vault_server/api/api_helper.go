package api

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
	"vault-auth-plugin/vault_server/dao/audit"
	"vault-auth-plugin/vault_server/models/audit"
)

var secretKey = []byte("secretKey")

func VerifyToken(r *http.Request) (bool, error) {
	if r.Header["Authorization"] != nil {
		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err == nil && token.Valid {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, errors.New("no token provided")
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func WriteLog(command string, r *http.Request) {
	audit := audit.Log{
		Time:    time.Now(),
		Ip:      r.RemoteAddr,
		Command: command,
	}
	auditdao.Insert(&audit)
}
