package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/users/models"

type UserDao interface {
	GetAll() []models.User
	CreateUser(user *models.User) *models.User
}
