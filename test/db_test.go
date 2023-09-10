package test

import (
	"log"
	"os"
	"testing"

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
func TestDatabaseConnection(t *testing.T) {
	os.Setenv("DATABASE", "root@tcp(localhost:3306)/db_user")

	Connect()

	t.Logf("Koneksi database berhasil!")

}
