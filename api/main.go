package main

import (
	"fmt"
	"github.com/fedehsq/api/api/admin"
	"github.com/fedehsq/api/api/log"
	"github.com/fedehsq/api/api/user"
	"github.com/fedehsq/api/config"
	"github.com/fedehsq/api/db"
	_ "github.com/fedehsq/api/docs"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"strings"
	"time"
)

// @title           Swagger Vault support API
// @version         1.0
// @description     This is an API Vault server support.

// @contact.name   API Support

// @host      localhost:19090
// @BasePath  /api/v1

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

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
		config.Conf.DbPassword,
	)
	if err != nil {
		log.Fatal(err)
	}
	//err = api.GetKey(config.Conf.ApiVaultToken)
	//if err != nil {
	//	log.Fatal(err)
	//}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/logs", logapi.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/admin-signin", adminapi.Signin).Methods("POST")
	r.HandleFunc("/api/v1/signup", userapi.Signup).Methods("POST")
	r.HandleFunc("/api/v1/signin", userapi.Signin).Methods("POST")
	r.HandleFunc("/api/v1/users", userapi.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/user", userapi.GetByUsername).Methods("GET")
	r.HandleFunc("/api/v1/user", userapi.Update).Methods("PUT")
	r.HandleFunc("/api/v1/user", userapi.Delete).Methods("DELETE")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	handler := cors.AllowAll().Handler(r)
	srv := &http.Server{
		Handler:      handler,
		Addr:         strings.Replace(config.Conf.ApiAddress, "http://", "", 1),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Vault server started at %s\n", config.Conf.ApiAddress)
	log.Fatal(srv.ListenAndServe())
}
