package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"gorm.io/gorm"
)

type CategoryDao struct {
	Db *gorm.DB
}

func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{Db: db}
}

func (d CategoryDao) GetByIds(ids []uint) ([]models.Category, error) {
	var categories []models.Category

	if err := d.Db.Find(&categories, "id in (?)", ids).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (d CategoryDao) GetAll() ([]models.Category, error) {
	var categories []models.Category

	if err := d.Db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (d CategoryDao) GetById(id uint) (*models.Category, error) {
	var category models.Category

	if err := d.Db.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (d CategoryDao) Create(category models.Category) (*models.Category, error) {
	if err := d.Db.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (d CategoryDao) Delete(category models.Category) error {
	if err := d.Db.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
