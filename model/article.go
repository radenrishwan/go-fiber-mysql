package model

type Article struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
