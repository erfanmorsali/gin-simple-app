package services

import (
	"errors"
	"fmt"
	"github.com/erfanmorsali/gin-simple-app/database/models"
	"github.com/erfanmorsali/gin-simple-app/users/dtos"
	"github.com/erfanmorsali/gin-simple-app/users/interfaces"
	"gorm.io/gorm"
)

type userService struct {
	userDao interfaces.UserDao
}

func NewUserService(userDao interfaces.UserDao) interfaces.UserService {
	return &userService{userDao: userDao}
}

func (s userService) GetAll() ([]dtos.UserOut, error) {
	var result []dtos.UserOut

	users, err := s.userDao.GetAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userOut := dtos.CreateUserOut(user)
		result = append(result, userOut)
	}

	return result, nil
}

func (s userService) Create(userIn dtos.UserIn) (*dtos.UserOut, error) {
	var user = models.User{Username: userIn.Username, Password: userIn.Password}

	createdUser, err := s.userDao.Create(&user)
	if err != nil {
		return nil, err
	}

	return &dtos.UserOut{ID: createdUser.ID, Username: createdUser.Username}, nil
}

func (s userService) GetById(id uint) (*dtos.UserOut, error) {
	user, err := s.userDao.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage := fmt.Sprintf("record with id : %d not found", id)
			return nil, errors.New(errorMessage)
		}
		return nil, err
	}
	userOut := dtos.CreateUserOut(*user)

	return &userOut, nil
}

func (s userService) DeleteById(id uint) error {
	user, err := s.userDao.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorMessage := fmt.Sprintf("record with id : %d not found", id)
			return errors.New(errorMessage)
		}
		return err
	}

	return s.userDao.Delete(user)
}
