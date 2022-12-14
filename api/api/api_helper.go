package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fedehsq/api/config"
	logdao "github.com/fedehsq/api/dao/log"
	"github.com/fedehsq/api/models/log"

	"github.com/golang-jwt/jwt"
)

type Key struct {
	Data struct {
		Data struct {
			Key string `json:"key"`
		} `json:"data"`
	} `json:"data"`
}

var secretKey []byte

// Get the secret key for generate and check JWT from the vault of the bastion host
func GetKey(token string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/secret/data/api", config.Conf.VaultAddress), nil)
	if err != nil {
		return err
	}
	body, err := doRequest(req, token)
	if err != nil {
		return err
	}
	key := Key{}
	err = json.Unmarshal(body, &key)
	if err != nil {
		return err
	}
	secretKey = []byte(key.Data.Data.Key)
	return nil
}

func doRequest(req *http.Request, vaultToken string) ([]byte, error) {
	req.Header.Set("X-Vault-Token", vaultToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}

func VerifyToken(r *http.Request) (bool, string, error) {
	if r.Header["Authorization"] != nil {
		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			return true, claims["identity"].(string), nil
		} else {
			return false, "", err
		}
	}
	return false, "", errors.New("no token provided")
}

func GenerateJWT(identity string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func WriteLog(method string, route string, callerIdentity string, ip string, body string) {
	// Save time as string
	audit := log.Log{
		Time:           time.Now().Format(time.UnixDate),
		Ip:             ip,
		CallerIdentity: callerIdentity,
		Method:         method,
		Route:          route,
		Body:           body,
	}
	logdao.Insert(&audit)
}
