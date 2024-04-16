package main

import (
	"github.com/joho/godotenv"
	"log"
	"pointsCounterCRUD/database"
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
}

func serveApplication() {
	routes := endpoints.NewRouts(database.DB)
	routes.AddPath()
	routes.Start(2137)
}
