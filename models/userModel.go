package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email           string `gorm:"unique"`
	Password        string
	EmailPromotions bool `gorm:"default:false"`
}
