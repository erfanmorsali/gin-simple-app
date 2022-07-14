package services

import (
	"github.com/erfanmorsali/gin-simple-app.git/database/models"
	"github.com/erfanmorsali/gin-simple-app.git/posts/dtos"
	"github.com/erfanmorsali/gin-simple-app.git/posts/interfaces"
	userInterfaces "github.com/erfanmorsali/gin-simple-app.git/users/interfaces"
)

type PostService struct {
	PostDao     interfaces.PostDao
	UserDao     userInterfaces.UserDao
	CategoryDao interfaces.CategoryDao
}

func NewPostService(dao interfaces.PostDao, userDao userInterfaces.UserDao, categoryDao interfaces.CategoryDao) *PostService {
	return &PostService{PostDao: dao, UserDao: userDao, CategoryDao: categoryDao}
}

func (p PostService) GetAllPosts() ([]dtos.PostOut, error) {
	var result []dtos.PostOut

	posts, err := p.PostDao.GetAll()
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		var postOut = dtos.CreatePostOut(*post)
		result = append(result, postOut)
	}

	return result, nil
}

func (p PostService) GetPostById(id uint) (*dtos.PostOut, error) {
	post, err := p.PostDao.GetById(id)
	if err != nil {
		return nil, err
	}

	postOut := dtos.CreatePostOut(*post)
	return &postOut, nil

}

func (p PostService) CreatePost(postIn dtos.PostIn) (*dtos.PostOut, error) {
	user, err := p.UserDao.GetById(postIn.UserID)
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	if postIn.CategoryIDS != nil && len(postIn.CategoryIDS) >= 1 {
		categories, err = p.CategoryDao.GetCategoriesByIds(postIn.CategoryIDS)
		if err != nil {
			return nil, err
		}
	}

	post := models.Post{User: *user, Title: postIn.Title, Description: postIn.Description, Categories: categories}
	createdPost, err := p.PostDao.CreatePost(post)
	if err != nil {
		return nil, err
	}

	postOut := dtos.CreatePostOut(*createdPost)
	return &postOut, nil
}

func (p PostService) DeletePostById(id uint) error {
	post, err := p.PostDao.GetById(id)
	if err != nil {
		return err
	}

	err = p.PostDao.DeletePost(*post)
	if err != nil {
		return err
	}

	return nil
}
