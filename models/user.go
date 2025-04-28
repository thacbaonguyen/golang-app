package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:100;not null; unique" json:"username"`
	Email     string         `gorm:"size:100;not null; unique" json:"email"`
	Password  string         `gorm:"size:100;not null;" json:"-"`
	FullName  string         `gorm:"size:100" json:"full_name"`
	RoleId    uint           `json:"role_id"`
	Role      Role           `gorm:"foreignKey:RoleId" json:"role,omitempty"`
	Posts     []Post         `gorm:"foreignKey:UserId" json:"posts,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
