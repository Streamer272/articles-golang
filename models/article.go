package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Id 		uint64 `gorm:"primaryKey"`
	Title 	string
	Text 	string
}
