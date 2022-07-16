package dtos

import (
	"github.com/erfanmorsali/gin-simple-app/database/models"
)

type PostOut struct {
	ID            uint           `json:"id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	UserInfo      *userInfo      `json:"user"`
	CategoryInfos []categoryInfo `json:"categories"`
}

func CreatePostOut(post models.Post) PostOut {
	var categoryInfos []categoryInfo
	var user *userInfo

	if len(post.Categories) >= 1 {
		for _, category := range post.Categories {
			var categoryInfo = categoryInfo{ID: category.ID, Name: category.Name}
			categoryInfos = append(categoryInfos, categoryInfo)
		}

	}

	if post.User.ID != 0 {
		user = &userInfo{ID: post.User.ID, Username: post.User.Username}
	}

	return PostOut{ID: post.ID, Title: post.Title, Description: post.Description, CategoryInfos: categoryInfos, UserInfo: user}
}
