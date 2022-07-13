package services

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"github.com/erfanmorsali/gin-simple-app.git/users/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
)

type UserService struct {
	UserDao interfaces.UserDao
}

func NewUserService(userDao interfaces.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}

func (s UserService) GetAllUsers() []dtos.UserOut {
	var result []dtos.UserOut
	users := s.UserDao.GetAll()
	for _, user := range users {
		userOut := dtos.CreateUserOut(user)
		result = append(result, userOut)
	}
	return result
}

func (s UserService) CreateUser(userIn dtos.UserIn) *dtos.UserOut {
	var user = models.User{Username: userIn.Username, Password: userIn.Password}

	createdUser := s.UserDao.CreateUser(&user)
	return &dtos.UserOut{ID: createdUser.ID, Username: createdUser.Username}
}

func (s UserService) GetUserById(id uint) (*dtos.UserOut, error) {
	user, err := s.UserDao.GetById(id)
	if err != nil {
		return nil, err
	}
	userOut := dtos.CreateUserOut(*user)
	return &userOut, nil
}

func (s UserService) DeleteUserById(id uint) error {
	user, err := s.UserDao.GetById(id)
	if err != nil {
		return err
	}

	return s.UserDao.DeleteUser(user)
}
