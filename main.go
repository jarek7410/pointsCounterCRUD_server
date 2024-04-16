package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"pointsCounterCRUD/database"
	"pointsCounterCRUD/database/model"
	"pointsCounterCRUD/endpoints"
)

func main() {
	loadEnv()

	loadDatabase()

	serveApplication()

	//Repo := dbStats.NewSQLGORMRepository(ctx, db)
	//if err := Repo.Migrate(); err != nil {
	//	log.Fatalln(err)
	//}

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	database.InitDatabase()
	database.DB.AutoMigrate(&model.Role{})
	database.DB.AutoMigrate(&model.User{})
	seedData()
}
func seedData() {
	var roles = []model.Role{{Name: "admin", Description: "Administrator role"},
		{Name: "customer", Description: "Authenticated customer role"},
		{Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []model.User{
		{Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}

	database.DB.Save(&roles)
	database.DB.Save(&user)
}

func serveApplication() {
	routes := endpoints.NewRouts(database.DB)
	routes.AddPath()
	routes.Start(2137)
}
