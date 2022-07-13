package models

type Post struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	UserID      uint
	User        User
}
