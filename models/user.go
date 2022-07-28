package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string
	Username string `gorm:"unique_index;not null"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	IsStaff  bool   `gorm:"not null"`
	IsAdmin  bool   `gorm:"not null"`
}
