package services

import (
	"errors"
	"fmt"
	"github.com/erfanmorsali/gin-simple-app/database/models"
	"github.com/erfanmorsali/gin-simple-app/posts/dtos"
	"github.com/erfanmorsali/gin-simple-app/posts/interfaces"
	"gorm.io/gorm"
)

type categoryService struct {
	categoryDao interfaces.CategoryDao
}

func NewCategoryService(categoryDao interfaces.CategoryDao) interfaces.CategoryService {
	return &categoryService{categoryDao: categoryDao}
}

func (s categoryService) GetAll() ([]dtos.CategoryOut, error) {
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

func (s categoryService) GetById(id uint) (*dtos.CategoryOut, error) {
	category, err := s.categoryDao.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage := fmt.Sprintf("record with id : %d not found", id)
			return nil, errors.New(errorMessage)
		}
		return nil, err
	}

	result := dtos.CreateCategoryOut(*category)

	return &result, nil
}

func (s categoryService) Create(categoryIn dtos.CategoryIn) (*dtos.CategoryOut, error) {
	category := models.Category{Name: categoryIn.Name}

	createdCategory, err := s.categoryDao.Create(category)
	if err != nil {
		return nil, err
	}

	categoryOut := dtos.CreateCategoryOut(*createdCategory)
	return &categoryOut, nil
}

func (s categoryService) DeleteById(id uint) error {
	category, err := s.categoryDao.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage := fmt.Sprintf("record with id : %d not found", id)
			return errors.New(errorMessage)
		}
		return err
	}

	err = s.categoryDao.Delete(*category)
	if err != nil {
		return err
	}

	return nil
}
