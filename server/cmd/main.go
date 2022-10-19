package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	admin "vault-auth-plugin/server/api/admin"
	user "vault-auth-plugin/server/api/user"
	sqldb "vault-auth-plugin/server/db"

	"github.com/gorilla/mux"
)

func main() {
	sqldb.InitDb()
	r := mux.NewRouter()
	r.HandleFunc("/admin-signin", admin.Signin).Methods("POST")

	r.HandleFunc("/signup", user.Signup).Methods("POST")
	r.HandleFunc("/signin", user.Signin).Methods("POST")
	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/user", user.GetUser).Methods("GET")
	r.HandleFunc("/user", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", user.DeleteUser).Methods("DELETE")
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
