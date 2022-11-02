package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"vault-auth-plugin/bh/config"
	"vault-auth-plugin/bh/user"

	"github.com/gorilla/mux"
)

func main() {
	err := config.LoadConfig("./bh")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/signin", userapi.Signin).Methods("POST")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Vault server started at %s\n", "http://127.0.0.1:5000")
	log.Fatal(srv.ListenAndServe())
}
