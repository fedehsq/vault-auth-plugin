package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	//"github.com/fedehsq/api/api"
	adminapi "github.com/fedehsq/api/api/admin"
	logapi "github.com/fedehsq/api/api/log"
	userapi "github.com/fedehsq/api/api/user"
	"github.com/fedehsq/api/config"
	"github.com/fedehsq/api/db"

	"github.com/gorilla/mux"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:19090
// @BasePath  /api/v1

//@securityDefinitions.apikey JWT
//@in header
//@name JWT

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
	srv := &http.Server{
		Handler:      r,
		Addr:         strings.Replace(config.Conf.ApiAddress, "http://", "", 1),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Vault server started at %s\n", config.Conf.ApiAddress)
	log.Fatal(srv.ListenAndServe())
}
