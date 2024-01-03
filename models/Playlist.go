package models

import (
	"gorm.io/gorm"
)

type Playlist struct {
	gorm.Model
	ID                          uint   `gorm:"primarykey" json:"id"`
	UserID                      uint   `json:"user_id,omitempty"`
	SongIds                     []int  `gorm:"type:int" json:"song_ids"`
	Title                       string `json:"title"`
	Description                 string `json:"description"`
	SongCount                   uint   `json:"song_count"`
	TotalTrackDurationInSeconds string `json:"total_track_duration_in_seconds,omitempty"`
}
