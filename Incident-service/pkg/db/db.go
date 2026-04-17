package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("postgres://postgres:Kwidrxl@1432@db.leznplzzvcfuqbtcghem.supabase.co:5432/postgres?sslmode=require")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
	}

	return db
}
