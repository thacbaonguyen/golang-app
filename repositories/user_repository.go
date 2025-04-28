package repositories

import "go-ginapp/models"

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id uint) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindByEmail(email string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	UpdatePassword(user models.User) error
	Delete(id uint) error
}
