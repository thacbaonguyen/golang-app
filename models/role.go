package models

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;size:100;not null;unique" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
}
