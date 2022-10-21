package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"vault-auth-plugin/server/api/admin"
	"vault-auth-plugin/server/api/user"
	"vault-auth-plugin/server/api/log"
	"vault-auth-plugin/server/db"

	"github.com/gorilla/mux"
)

func main() {
	sqldb.InitDb()
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
		Handler: r,
		Addr:    "127.0.0.1:19090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Server started at port 19090")
	log.Fatal(srv.ListenAndServe())
}
