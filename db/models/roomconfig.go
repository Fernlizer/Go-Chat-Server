package models

import "gorm.io/gorm"

type RoomConfig struct {
	gorm.Model
	RoomID      int    `gorm:"default:NULL"`
	Password    string `gorm:"default:NULL"`
	AllowUserID int    `gorm:"default:NULL"`
}
