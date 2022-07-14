package dtos

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
)

type UserOut struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Posts    []postInfo
}

func CreateUserOut(user models.User) UserOut {
	var postInfos []postInfo

	if len(user.Posts) >= 1 {
		for _, post := range user.Posts {
			postInfo := postInfo{ID: post.ID, Title: post.Title, Description: post.Description}
			postInfos = append(postInfos, postInfo)
		}
	}

	return UserOut{ID: user.ID, Username: user.Username, Posts: postInfos}
}
