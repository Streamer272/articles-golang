package models

type Article struct {
	Id     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserId uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserId;references:id" json:"-"`
}
