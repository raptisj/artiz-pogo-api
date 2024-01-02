package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	ID       uint   `gorm:"primarykey" json:"id"`
	ArtistID uint   `json:"artist_id"`
	Title    string `json:"title"`
	Album    string `json:"album"`
	Year     int    `json:"year"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
}
