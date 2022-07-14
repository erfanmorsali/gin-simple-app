package models

type Category struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Posts []Post `gorm:"many2many:posts_categories"`
}
