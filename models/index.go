package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create table in database and seed tables with data from json file
	if err = db.AutoMigrate(&Artist{}); err == nil && db.Migrator().HasTable(&Artist{}) {
		if err := db.First(&Artist{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			SeedArtists(db)
		}
	}

	if err = db.AutoMigrate(&Song{}); err == nil && db.Migrator().HasTable(&Song{}) {
		if err := db.First(&Song{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			SeedSongs(db)
		}
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&LikedSong{})
	db.AutoMigrate(&FollowArtist{})
	fmt.Println("Database is ready")

	DB = db
}

func SeedArtists(db *gorm.DB) {
	data, err := os.ReadFile("artists.json")
	if err != nil {
		panic(err)
	}

	artists := []Artist{}
	err = json.Unmarshal(data, &artists)
	if err != nil {
		panic(err)
	}

	for _, artist := range artists {
		db.Create(&artist)
	}

	fmt.Println("Artists seeded")
}

func SeedSongs(db *gorm.DB) {
	data, err := os.ReadFile("songs.json")
	if err != nil {
		panic(err)
	}

	songs := []Song{}
	err = json.Unmarshal(data, &songs)
	if err != nil {
		panic(err)
	}

	for _, song := range songs {
		db.Create(&song)
	}

	fmt.Println("Songs seeded")
}
