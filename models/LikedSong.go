package models

import (
	"gorm.io/gorm"
)

type LikedSong struct {
	gorm.Model
	ID     uint `gorm:"primarykey" json:"id"`
	UserID uint `json:"user_id"`
	SongID uint `json:"song_id"`
}
