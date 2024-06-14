package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	TypeID      int        `gorm:"not null"`
	ConfigID    int        `gorm:"not null"`
	Password    string     `gorm:"not null"`
	IsActive    bool       `gorm:"not null"`
	Type        RoomType   `gorm:"foreignKey:TypeID"`
	Config      RoomConfig `gorm:"foreignKey:ConfigID"`
}
