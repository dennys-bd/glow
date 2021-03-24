package main

import (
	"log"
	"time"

	"github.com/dennys-bd/glow/ops/migrate/migrations"
	"github.com/dennys-bd/glow/repository"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = repository.GetDB()
}

func initSchema(tx *gorm.DB) error {
	type Class struct {
		ID        uint
		Capacity  uint
		Name      string
		StartDate time.Time
		EndDate   time.Time
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`
	}

	return tx.AutoMigrate(&Class{})
}

func main() {
	log.Print("Starting database migration...")

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.GetMigrations())

	m.InitSchema(initSchema)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Print("Migration did run successfully!")
}
