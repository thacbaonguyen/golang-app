package models

import "time"

type Post struct {
	ID        uint      `gorm:"size:100;not null;unique" json:"id"`
	Title     string    `gorm:"size:255;not null;" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	UserId    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserId" json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
