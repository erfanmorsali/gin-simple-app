package dtos

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
)

type UserOut struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Posts    []PostInfo
}

type PostInfo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateUserOut(user models.User) UserOut {
	var postInfos []PostInfo

	if len(user.Posts) >= 1 {
		for _, post := range user.Posts {
			postInfo := PostInfo{ID: post.ID, Title: post.Title, Description: post.Description}
			postInfos = append(postInfos, postInfo)
		}
	}

	return UserOut{ID: user.ID, Username: user.Username, Posts: postInfos}
}
