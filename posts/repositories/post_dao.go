package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"gorm.io/gorm"
)

type postDao struct {
	db *gorm.DB
}

func NewPostDao(db *gorm.DB) *postDao {
	return &postDao{db: db}
}

func (d postDao) GetAll() ([]*models.Post, error) {
	var posts []*models.Post

	if err := d.db.Preload("Categories").Joins("User").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (d postDao) GetById(id uint) (*models.Post, error) {
	var post models.Post

	if err := d.db.Joins("User").Preload("Categories").First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (d postDao) Create(post models.Post) (*models.Post, error) {
	if err := d.db.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (d postDao) Delete(post models.Post) error {
	if err := d.db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
