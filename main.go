package main

import (
	catc "MeowGoWithDB/services/cat/controller"
	cats "MeowGoWithDB/services/cat/service"

	_gorm "MeowGoWithDB/services/db/gorm"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()

	if err != nil {
		log.Print(err.Error())
	}
}

func main() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	dbhost := viper.GetString("mysql.host")
	dbport := viper.GetString("mysql.port")
	dbusername := viper.GetString("mysql.username")
	dbpassword := viper.GetString("mysql.password")
	dbname := viper.GetString("mysql.dbname")

	sqlConnection := _gorm.ConnectDB(dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname)

	defer func() {
		sql, err := sqlConnection.DB()
		if err != nil {
			log.Panic(err.Error())
		}
		sql.Close()
	}()

	cat := r.PathPrefix("/cats").Subrouter()
	catc.Controller(cat)
	cats.InitCatRepository(sqlConnection)

	r.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"alive":true}`))
	}).Methods("GET")

	log.Print("running on :" + viper.GetString("port"))

	http.ListenAndServe(":"+viper.GetString("port"), handlers.CORS(header, methods, origins, credentials)(r))
}
