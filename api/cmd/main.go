package main

import (
	"fmt"
	"github.com/fedehsq/api/api"
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
// @BasePath  /api

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

	if config.Conf.Develop == 0 {
		err = api.GetKey(config.Conf.ApiVaultToken)
		if err != nil {
			log.Fatal(err)
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/log/get-all", logapi.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/admin/signin", adminapi.Signin).Methods("POST")
	r.HandleFunc("/api/v1/user/signup", userapi.Signup).Methods("POST")
	r.HandleFunc("/api/v1/user/signin", userapi.Signin).Methods("POST")
	r.HandleFunc("/api/v1/user/get-all", userapi.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/user/get", userapi.GetByUsername).Methods("GET")
	r.HandleFunc("/api/v1/user/update", userapi.Update).Methods("PUT")
	r.HandleFunc("/api/v1/user/delete", userapi.Delete).Methods("DELETE")
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
