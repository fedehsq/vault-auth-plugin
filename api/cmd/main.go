package main

import (
	"fmt"
	"github.com/fedehsq/vault-auth-plugin/api/api"
	"github.com/fedehsq/vault-auth-plugin/api/api/admin"
	"github.com/fedehsq/vault-auth-plugin/api/api/log"
	"github.com/fedehsq/vault-auth-plugin/api/api/user"
	"github.com/fedehsq/vault-auth-plugin/api/config"
	"github.com/fedehsq/vault-auth-plugin/api/db"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	err = db.InitDb(
		config.Conf.DbAddress,
		config.Conf.DbPort,
		config.Conf.DbUser,
		config.Conf.DbName,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = api.GetKey(config.Conf.ApiVaultToken)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/logs", logapi.GetAll).Methods("GET")
	r.HandleFunc("/admin-signin", adminapi.Signin).Methods("POST")
	r.HandleFunc("/signup", userapi.Signup).Methods("POST")
	r.HandleFunc("/signin", userapi.Signin).Methods("POST")
	r.HandleFunc("/users", userapi.GetAll).Methods("GET")
	r.HandleFunc("/user", userapi.GetByUsername).Methods("GET")
	r.HandleFunc("/user", userapi.Update).Methods("PUT")
	r.HandleFunc("/user", userapi.Delete).Methods("DELETE")
	srv := &http.Server{
		Handler:      r,
		Addr:         strings.Replace(config.Conf.ApiAddress, "http://", "", 1),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Vault server started at %s\n", config.Conf.ApiAddress)
	log.Fatal(srv.ListenAndServe())
}
