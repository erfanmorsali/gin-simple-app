package repositories

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d UserDao) GetAll() []models.User {

	var users []models.User
	d.db.Preload("Posts").Find(&users)

	return users
}

func (d UserDao) Create(user *models.User) (*models.User, error) {
	if err := d.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (d UserDao) GetById(id uint) (*models.User, error) {
	var user models.User

	if err := d.db.Preload("Posts").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d UserDao) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := d.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (d UserDao) Delete(user *models.User) error {
	return d.db.Delete(user).Error
}
