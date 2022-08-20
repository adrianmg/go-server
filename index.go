package main

import (
	"fmt"
	"log"
	"time"

	// "net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	loadEnv()
	db := getDatabase()

	ping, _ := db.DB()
	if err := ping.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ping.Stats().OpenConnections)

	var stats []Stat
	result := db.Find(&stats)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(stats)
}

type Stat struct {
	ID        int
	Slug      string
	Views     int
	Downloads int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
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
