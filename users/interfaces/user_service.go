package interfaces

import "github.com/erfanmorsali/gin-simple-app.git/users/models"

type UserService interface {
	GetAllUsers() []models.UserOut
}
