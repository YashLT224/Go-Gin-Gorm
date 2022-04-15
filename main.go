package main

import (
	"gin-gorm/configs"
	"gin-gorm/database"
	"gin-gorm/models"
	"gin-gorm/repositories"
	"log"
)

func main() {
	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Contact{})

	contactRepository := repositories.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8000")
}
