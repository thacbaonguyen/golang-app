package utils

import (
	"errors"
	"go-ginapp/models"
	"gorm.io/gorm"
)

func InitializeRoles(db *gorm.DB) error {
	defaultRoles := []models.Role{
		{Name: "user"},
		{Name: "admin"},
	}

	for _, role := range defaultRoles {
		var existingRole models.Role
		err := db.Where("name = ?", role.Name).First(&existingRole).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(role).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}
