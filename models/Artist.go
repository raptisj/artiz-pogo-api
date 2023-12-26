package models

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	ID         string `gorm:"primarykey"`
	Name       string `json:"name"`
	BirthYear  int    `json:"birth_year"`
	DeathYear  int    `json:"death_year"`
	Instrument string `json:"instrument"`
	Genre      string `json:"genre"`
	Bio        string `json:"bio"`
}
