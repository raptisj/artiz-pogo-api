package models

import "gorm.io/gorm"

type FollowArtist struct {
	gorm.Model
	ID       uint `gorm:"primarykey"`
	UserID   uint `json:"user_id"`
	ArtistID uint `json:"artist_id"`
}
