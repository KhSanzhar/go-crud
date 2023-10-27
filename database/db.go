package database

import (
	"api/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "Golang"
)

var (
	DB *gorm.DB
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)

func InitDB() {
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("Connection to database failed!")
	}
	err = db.AutoMigrate(&model.Item{})
	if err != nil {
		return
	}
	DB = db
}