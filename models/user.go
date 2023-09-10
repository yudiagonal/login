package models

type User struct {
	Id       uint64 `gorm:"primary_key;not_null" json:"id"`
	Name     string `gorm:"varchar(100)" json:"name"`
	Username string `gorm:"varchar(100)" json:"username"`
	Password string `gorm:"varchar(100)" json:"password"`
}
