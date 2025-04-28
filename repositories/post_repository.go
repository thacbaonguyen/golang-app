package repositories

import "go-ginapp/models"

type PostRepository interface {
	FindAll() ([]models.Post, error)
	FindById(id uint) (models.Post, error)
	FindByUser(userId uint) ([]models.Post, error)
	Create(post models.Post) (models.Post, error)
	Update(post models.Post) (models.Post, error)
	Delete(id uint) error
}
