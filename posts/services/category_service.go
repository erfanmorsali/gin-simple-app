package services

import "github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"

type CategoryService struct {
	CategoryDao interfaces.CategoryDao
}

func NewCategoryService(categoryDao interfaces.CategoryDao) *CategoryService {
	return &CategoryService{CategoryDao: categoryDao}
}
