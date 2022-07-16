package dtos

import "github.com/erfanmorsali/gin-simple-app/database/models"

type CategoryOut struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateCategoryOut(category models.Category) CategoryOut {
	return CategoryOut{ID: category.ID, Name: category.Name}
}
