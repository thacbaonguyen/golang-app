package repositories

import (
	"errors"
	"go-ginapp/models"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepositoryImpl(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (p *PostRepositoryImpl) FindAll() ([]models.Post, error) {
	//TODO implement me
	var posts []models.Post
	err := p.db.Preload("User").Find(&posts).Error
	return posts, err
}

func (p *PostRepositoryImpl) FindById(id uint) (models.Post, error) {
	var post models.Post
	err := p.db.Preload("User").Find(&post, id).Error
	return post, err
}

func (p *PostRepositoryImpl) FindByUser(userId uint) ([]models.Post, error) {
	var posts []models.Post
	err := p.db.Preload("User").Find(&posts).Where("userId = ?", userId).Error
	return posts, err
}

func (p *PostRepositoryImpl) Create(post models.Post) (models.Post, error) {
	//TODO implement me
	err := p.db.Create(post).Error
	return post, err
}

func (p *PostRepositoryImpl) Update(post models.Post) (models.Post, error) {
	//TODO implement me
	err := p.db.Save(post).Error
	return post, err
}

func (p *PostRepositoryImpl) Delete(id uint) error {
	//TODO implement me
	var post models.Post
	err := p.db.First(&post, id).Error
	if err != nil {
		return errors.New("post not found")
	}
	return p.db.Delete(post).Error
}
