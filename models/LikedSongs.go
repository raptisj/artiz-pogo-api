package models

import "gorm.io/gorm"

type LikedSongs struct {
	gorm.Model
	ID     uint   `gorm:"primarykey"`
	UserId []User `gorm:"foreignKey:ID"`
	SongId []Song `gorm:"foreignKey:ID"`
}
