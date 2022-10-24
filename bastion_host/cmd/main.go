package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"vault-auth-plugin/bastion_host/api/user"
)

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/signup", userapi.Signup).Methods("POST")
	r.HandleFunc("/signin", userapi.Signin).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:19091",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started at port 19091")
	log.Fatal(srv.ListenAndServe())
}
