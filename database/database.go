package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	log.Println("loading environments")
	var dbHost = os.Getenv("DB_HOST")
	var dbPost = os.Getenv("DB_PORT")
	var dbUser = os.Getenv("DB_USER")
	var dbName = os.Getenv("DB_NAME")
	var dbPass = os.Getenv("BD_PASSWORD")

	log.Println("Connecting to db")

	psqlconn :=
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPost, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, "fail to connect to db")
	}

	log.Println("`Successfully connected to the database")
	return db
}
