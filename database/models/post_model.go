package models

type Post struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	UserID      uint
	User        User
	Categories  []Category `gorm:"many2many:posts_categories"`
}
