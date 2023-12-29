package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       int    `gorm:"primarykey"`
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}
