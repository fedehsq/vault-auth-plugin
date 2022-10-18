package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	auth "vault-auth-plugin/server/api"
	sqldb "vault-auth-plugin/server/db"

	"github.com/gorilla/mux"
)

func main() {
	sqldb.InitDb()
	r := mux.NewRouter()
	r.HandleFunc("/signup", auth.Signup).Methods("POST")
	r.HandleFunc("/signin", auth.Signin).Methods("POST")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:19090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started at port 19090")
	log.Fatal(srv.ListenAndServe())
}
