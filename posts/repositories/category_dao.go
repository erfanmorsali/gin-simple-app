package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"
	"gorm.io/gorm"
)

type categoryDao struct {
	db *gorm.DB
}

func NewCategoryDao(db *gorm.DB) interfaces.CategoryDao {
	return &categoryDao{db: db}
}

func (d categoryDao) GetByIds(ids []uint) ([]models.Category, error) {
	var categories []models.Category

	if err := d.db.Find(&categories, "id in (?)", ids).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (d categoryDao) GetAll() ([]models.Category, error) {
	var categories []models.Category

	if err := d.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (d categoryDao) GetById(id uint) (*models.Category, error) {
	var category models.Category

	if err := d.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (d categoryDao) Create(category models.Category) (*models.Category, error) {
	if err := d.db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (d categoryDao) Delete(category models.Category) error {
	if err := d.db.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
