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

func (d UserDao) GetById(id uint) (*models.User, error) {
	var user models.User

	if err := d.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d UserDao) DeleteUser(user *models.User) error {
	return d.Db.Delete(user).Error

}
