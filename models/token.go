package models

import (
	"time"
)

type Token struct {
	Id      uint      `gorm:"primaryKey" json:"id"`
	Expires time.Time `json:"expires"`
	UserId  uint      `json:"user_id"`
	User    User      `gorm:"foreignKey:UserId;references:id" json:"-"`
}
