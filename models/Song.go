package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	ID        uint      `gorm:"primarykey"`
	ArtistID  uint      `json:"artist_id"`
	Title     string    `json:"title"`
	Album     string    `json:"album"`
	Year      int       `json:"year"`
	Duration  string    `json:"duration"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
