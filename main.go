package main

import (
	"fmt"
	"log"
	"login/controller"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root@tcp(127.0.0.1:3306)/db_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	Connect()
	fmt.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
