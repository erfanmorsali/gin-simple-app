package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/users/models"
	"gorm.io/gorm"
)

type UserDao struct {
	Db *gorm.DB
}

func newUserDao(db *gorm.DB) *UserDao {
	return &UserDao{Db: db}
}

func (d UserDao) GetAll() []models.User {

	var users []models.User
	d.Db.Find(&users)

	return users
}

func (d UserDao) CreateUser(user *models.User) *models.User {
	d.Db.Create(user)
	return user
}
