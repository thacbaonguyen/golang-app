package repositories

import (
	"go-ginapp/models"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{db}
}

func (r *RoleRepositoryImpl) FindAll() ([]models.Role, error) {
	//TODO implement me
	var roles []models.Role
	err := r.db.Find(&roles).Error
	return roles, err
}

func (r *RoleRepositoryImpl) FindByName(name string) (models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return role, err
}

func (r *RoleRepositoryImpl) CreateRole(role models.Role) (models.Role, error) {
	//TODO implement me
	err := r.db.Create(&role).Error
	return role, err
}

func (r *RoleRepositoryImpl) UpdateRole(role models.Role) (models.Role, error) {
	//TODO implement me
	err := r.db.Save(role).Error
	return role, err
}
