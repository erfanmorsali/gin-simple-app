package services

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"github.com/erfanmorsali/gin-simple-app.git/posts/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"
)

type CategoryService struct {
	categoryDao interfaces.CategoryDao
}

func NewCategoryService(categoryDao interfaces.CategoryDao) *CategoryService {
	return &CategoryService{categoryDao: categoryDao}
}

func (s CategoryService) GetAll() ([]dtos.CategoryOut, error) {
	categories, err := s.categoryDao.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dtos.CategoryOut

	for _, category := range categories {
		categoryOut := dtos.CreateCategoryOut(category)
		result = append(result, categoryOut)
	}

	return result, nil
}

func (s CategoryService) GetById(id uint) (*dtos.CategoryOut, error) {
	category, err := s.categoryDao.GetById(id)
	if err != nil {
		return nil, err
	}

	result := dtos.CreateCategoryOut(*category)

	return &result, nil
}

func (s CategoryService) Create(categoryIn dtos.CategoryIn) (*dtos.CategoryOut, error) {
	category := models.Category{Name: categoryIn.Name}

	createdCategory, err := s.categoryDao.Create(category)
	if err != nil {
		return nil, err
	}

	categoryOut := dtos.CreateCategoryOut(*createdCategory)
	return &categoryOut, nil
}

func (s CategoryService) DeleteById(id uint) error {
	category, err := s.categoryDao.GetById(id)
	if err != nil {
		return err
	}

	err = s.categoryDao.Delete(*category)
	if err != nil {
		return err
	}

	return nil
}
