package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "main/pkg/config"
	domain "main/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	fmt.Println("====Connecting Database")
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.User{},&domain.Admin{})

	return db, dbErr
}
