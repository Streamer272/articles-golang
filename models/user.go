package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id 			uint64 `gorm:"primaryKey"`
	Username 	string `gorm:"unique"`
	Email		string `gorm:"unique"`
	Password 	string
}
