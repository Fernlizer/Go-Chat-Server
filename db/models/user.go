// db/models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `gorm:"default:NULL" validate:"required"`
	LastName    string `gorm:"default:NULL" validate:"required"`
	DisplayName string `gorm:"default:NULL"`
	Email       string `gorm:"default:NULL" validate:"required,email,uniqueEmail"`
	Username    string `gorm:"default:NULL" validate:"required,min=3,uniqueUsername"`
	Password    string `gorm:"default:NULL" validate:"required,min=6"`
	Role        int    `gorm:"default:NULL"`
	IsActive    bool   `gorm:"default:NULL"`
}
