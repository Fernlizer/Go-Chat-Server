package models

import "gorm.io/gorm"

type RoomType struct {
	gorm.Model
	Name        string `gorm:"default:NULL"`
	Description string `gorm:"default:NULL"`
	CreatedBy   int    `gorm:"default:NULL"`
	IsActive    bool   `gorm:"default:NULL"`
}
