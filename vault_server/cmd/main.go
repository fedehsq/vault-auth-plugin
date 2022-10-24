package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"vault-auth-plugin/config"
	"vault-auth-plugin/vault_server/api/admin"
	"vault-auth-plugin/vault_server/api/log"
	"vault-auth-plugin/vault_server/api/user"
	"vault-auth-plugin/vault_server/db"

	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	sqldb.InitDb(
		config.Conf.VaultServerDbAddress,
		config.Conf.VaultServerDbPort,
		config.Conf.VaultServerDbUser,
		config.Conf.VaultServerDbName,
	)
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
		Addr:         strings.Replace(config.Conf.VaultServerAddress, "http://", "", 1),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Vault server started at %s\n", config.Conf.VaultServerAddress)
	log.Fatal(srv.ListenAndServe())
}
