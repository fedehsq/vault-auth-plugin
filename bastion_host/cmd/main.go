package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"vault-auth-plugin/bastion_host/api/user"
	"vault-auth-plugin/config"

	"github.com/gorilla/mux"
)

func main() {
	// Load config
	err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/signin", userapi.Signin).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         strings.Replace(config.Conf.BastionHostAddress, "http://", "", 1),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Bastion host started at %s\n", config.Conf.BastionHostAddress)
	log.Fatal(srv.ListenAndServe())
}
