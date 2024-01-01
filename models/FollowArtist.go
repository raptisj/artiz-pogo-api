package models

import "gorm.io/gorm"

type FollowArtist struct {
	gorm.Model
	ID       uint     `gorm:"primarykey"`
	UserId   []User   `gorm:"foreignKey:ID"`
	ArtistId []Artist `gorm:"foreignKey:ID"`
}
