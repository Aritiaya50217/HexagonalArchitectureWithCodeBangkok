package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Aritiaya50217/CodeBangkok/handler"
	"github.com/Aritiaya50217/CodeBangkok/repository"
	"github.com/Aritiaya50217/CodeBangkok/service"
	"github.com/dvln/viper"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	initConfig()
	initTimeZone()
	db := initDataBase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods("GET")

	log.Printf("Banking service started at port : %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig() // อ่านไฟล์ config
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDataBase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime-true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
