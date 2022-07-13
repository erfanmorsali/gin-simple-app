package dtos

import "github.com/erfanmorsali/gin-simple-app.git/users/dtos"

type PostOut struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	UserInfo    dtos.UserInfo `json:"user"`
}
