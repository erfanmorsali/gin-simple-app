package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"gorm.io/gorm"
)

type PostDao struct {
	Db *gorm.DB
}

func NewPostDao(db *gorm.DB) *PostDao {
	return &PostDao{Db: db}
}

func (d PostDao) GetAll() ([]*models.Post, error) {
	var posts []*models.Post

	if err := d.Db.Preload("Categories").Joins("User").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (d PostDao) GetById(id uint) (*models.Post, error) {
	var post models.Post

	if err := d.Db.Joins("User").Preload("Categories").First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (d PostDao) CreatePost(post models.Post) (*models.Post, error) {
	if err := d.Db.Create(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (d PostDao) DeletePost(post models.Post) error {
	if err := d.Db.Delete(&post).Error; err != nil {
		return err
	}
	return nil
}
