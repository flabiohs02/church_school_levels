package domain

import (
	"gorm.io/gorm"
)

type SchoolLevel struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100);unique;not null"`
	Description string `gorm:"type:varchar(255)"`
	IsActive    bool   `gorm:"default:true"`
}
