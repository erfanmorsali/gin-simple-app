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

func (s UserService) GetAll() []dtos.UserOut {
	var result []dtos.UserOut

	users := s.UserDao.GetAll()
	for _, user := range users {
		userOut := dtos.CreateUserOut(user)
		result = append(result, userOut)
	}

	return result
}

func (s UserService) Create(userIn dtos.UserIn) (*dtos.UserOut, error) {
	var user = models.User{Username: userIn.Username, Password: userIn.Password}

	createdUser, err := s.UserDao.Create(&user)
	if err != nil {
		return nil, err
	}

	return &dtos.UserOut{ID: createdUser.ID, Username: createdUser.Username}, nil
}

func (s UserService) GetById(id uint) (*dtos.UserOut, error) {
	user, err := s.UserDao.GetById(id)
	if err != nil {
		return nil, err
	}
	userOut := dtos.CreateUserOut(*user)

	return &userOut, nil
}

func (s UserService) DeleteById(id uint) error {
	user, err := s.UserDao.GetById(id)
	if err != nil {
		return err
	}

	return s.UserDao.Delete(user)
}
