package repository

import (
	"log"
	"os"

	"github.com/dennys-bd/glow/ops/projectpath"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	var err error
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else if os.Getenv("ENVIRONMENT") == "TEST" {
		dir := projectpath.Root()
		db, _ = gorm.Open(sqlite.Open(dir+"/test.db"), &gorm.Config{})
		return db
	}
	db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}

func init() {
	db = GetDB()
}
