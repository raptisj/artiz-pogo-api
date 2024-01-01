package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
