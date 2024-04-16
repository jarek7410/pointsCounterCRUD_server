package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"pointsCounterCRUD/database"
	"pointsCounterCRUD/endpoints"
	model2 "pointsCounterCRUD/model"
)

func main() {
	loadEnv()

	repo := loadDatabase()

	serveApplication(repo)

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() *database.Repo {
	r := database.InitDatabase()
	err := Migrate(r)
	if err != nil {
		log.Fatalln("database migration do not work")
	}

	seedData(r)

	return r
}
func seedData(r *database.Repo) {
	var roles = []model2.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []model2.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
	r.DB.Save(&roles)
	r.DB.Save(&user)
}
func Migrate(r *database.Repo) error {
	err := r.DB.AutoMigrate(&model2.Role{}, &model2.User{})
	if err != nil {
		return err
	}
	return nil
}

func serveApplication(repo *database.Repo) {
	routes := endpoints.NewRouts(repo)
	routes.AddPaths()
	routes.AddAuthPaths()
	routes.ServeApplication()

	routes.Start(2137)
	fmt.Println("Server running on port 8000")
}
