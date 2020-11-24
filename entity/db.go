package entity

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// import for postgres sql driver used by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB : Reference for DB
var DB *gorm.DB

var err error

// Init : initializes the database and store it's reference in DB
func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=1995 dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	return DB
}

// GetDB : Using this function to get a connection.
func GetDB() *gorm.DB {
	return DB
}
