package main

import (
	"fmt"
	"log"

	// "net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stat struct {
	id        int
	slug      string
	views     int
	downloads int
	createdAt time.Time
	updatedAt time.Time
}

func main() {
	loadEnv()
	db := getDatabase()

	var stats []Stat
	result := db.Find(&stats)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println(stats)
}

func getDatabase() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	return db
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
