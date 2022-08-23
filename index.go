package main

import (
	"encoding/json"
	"log"
	"time"

	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	loadEnv()

	db := getDatabase()

	http.HandleFunc("/", GetStats(db))

	http.ListenAndServe(":8080", nil)
}

func GetStats(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var stats []Stat
		db.Find(&stats)

		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		enc.Encode(stats)
	}
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
		log.Printf("Error loading .env file: %s", err)
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
