package main

import (
	"log"
	"main/pkg/config"
	di "main/pkg/di"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

	// //db:=db.ConnectDatabase()
	// db.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	// db.Db.AutoMigrate(&models.User{},&models.Admin{},&models.Product{},&models.Category{})
	// //DB.Db.Model(&models.Product{}).BelongsTo(&models.Category{}, "Category_ID")
}
