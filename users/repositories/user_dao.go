package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
	"gorm.io/gorm"
)

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) interfaces.UserDao {
	return &userDao{db: db}
}

func (d userDao) GetAll() ([]models.User, error) {

	var users []models.User
	if err := d.db.Preload("Posts").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (d userDao) Create(user *models.User) (*models.User, error) {
	if err := d.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d userDao) GetById(id uint) (*models.User, error) {
	var user models.User

	if err := d.db.Preload("Posts").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d userDao) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := d.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (d userDao) Delete(user *models.User) error {
	return d.db.Delete(user).Error
}
