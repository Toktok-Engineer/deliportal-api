package config

import (
	"deliportal-api/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}

	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s ", host, user, dbName, password, dbPort)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	db.AutoMigrate()
	db.AutoMigrate(&model.Vehicle{})
	db.AutoMigrate(&model.Asuransi{})
	db.AutoMigrate(&model.AsuransiRekening{})
	db.AutoMigrate(&model.STNKReminder{})
	db.AutoMigrate(&model.AsuransiReminder{})
	db.AutoMigrate(&model.KIRReminder{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbc, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbc.Close()
}
