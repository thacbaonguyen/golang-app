package repositories

import "go-ginapp/models"

type RoleRepository interface {
	FindAll() ([]models.Role, error)
	FindByName(name string) (models.Role, error)
	CreateRole(role models.Role) (models.Role, error)
	UpdateRole(role models.Role) (models.Role, error)
}
