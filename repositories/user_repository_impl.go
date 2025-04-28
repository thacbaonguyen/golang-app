package repositories

import (
	"errors"
	"go-ginapp/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindAll() ([]models.User, error) {
	//TODO implement me
	var users []models.User
	err := r.db.Preload("Role").Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) FindById(id uint) (models.User, error) {
	//TODO implement me
	var user models.User
	err := r.db.Preload("Role").First(&user, id).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	//TODO implement me
	var user models.User
	err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	//TODO implement me
	var user models.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Create(user models.User) (models.User, error) {
	//TODO implement me
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.Password = string(hashedPassword)
	err = r.db.Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(user models.User) (models.User, error) {
	//TODO implement me
	err := r.db.Save(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) UpdatePassword(user models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint) error {
	//TODO implement me
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return errors.New("user not found")
	}

	return r.db.Delete(&user).Error
}
